package app

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/tuxiedev/gotweet/pkg/outputs/console"
	"github.com/tuxiedev/gotweet/pkg/structs"
)

// RunApp starts the main loop of the app
func RunApp(twitterConfig structs.TwitterCredentials, outputName string, outputConfig interface{}) {

	tweets := make(chan *twitter.Tweet)

	config := oauth1.NewConfig(twitterConfig.ConsumerKey, twitterConfig.ConsumerSecret)
	token := oauth1.NewToken(twitterConfig.APIKey, twitterConfig.APISecret)

	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter Client
	client := twitter.NewClient(httpClient)

	// Convenience Demux demultiplexed stream messages
	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {
		tweets <- tweet
	}
	demux.DM = func(dm *twitter.DirectMessage) {
		fmt.Println(dm.SenderID)
	}
	demux.Event = func(event *twitter.Event) {
		fmt.Printf("%#v\n", event)
	}

	filterParams := &twitter.StreamFilterParams{
		Track:         []string{"iphone"},
		StallWarnings: twitter.Bool(true),
	}
	println("Creating the stream")

	stream, err := client.Streams.Filter(filterParams)
	if err != nil {
		log.Fatal(err)
	}

	go demux.HandleChan(stream.Messages)

	if outputName == "console" {
		go console.Output(tweets)
	}

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)

	fmt.Println("Stopping Stream...")
	stream.Stop()
	close(tweets)

}
