package app

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/tuxiedev/gotweet/pkg/outputs"
	"github.com/tuxiedev/gotweet/pkg/structs"
)

// RunApp starts the main loop of the app
func RunApp(twitterCredentials structs.TwitterCredentials, outputName string, outputConfig interface{}) {

	tweets := make(chan *twitter.Tweet)

	config := oauth1.NewConfig(twitterCredentials.ConsumerKey, twitterCredentials.ConsumerSecret)
	token := oauth1.NewToken(twitterCredentials.APIKey, twitterCredentials.APISecret)

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
		Track:         []string{"iphone", "cat", "dog"},
		StallWarnings: twitter.Bool(true),
	}
	println("Creating the stream")

	stream, err := client.Streams.Filter(filterParams)
	if err != nil {
		log.Fatal(err)
	}

	go demux.HandleChan(stream.Messages)

	// configure and start output
	var output outputs.Output
	switch outputName {
	case "console":
		output = &outputs.Console{}
	case "kafka":
		output = &outputs.Kafka{Config: outputConfig}
	}
	output.Init()
	go output.Start(tweets)

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)

	fmt.Println("Stopping Stream...")
	output.Stop()
	stream.Stop()
	close(tweets)

}
