package twitter

import (
	"log"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/tuxiedev/gotweet/pkg/structs"
)

//StartConsumingFromTwitter starts the twitter stream consumption
func StartConsumingFromTwitter(t structs.TwitterConfig, tweets chan *twitter.Tweet) (*twitter.Stream, error) {
	config := oauth1.NewConfig(t.Credentials.APIKey, t.Credentials.APISecret)
	token := oauth1.NewToken(t.Credentials.AccessToken, t.Credentials.AccessSecret)

	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter Client
	client := twitter.NewClient(httpClient)

	// Convenience Demux demultiplexed stream messages
	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {
		tweets <- tweet
	}
	demux.DM = func(dm *twitter.DirectMessage) {
		log.Println(dm.SenderID)
	}
	demux.Event = func(event *twitter.Event) {
		log.Printf("%v\n", event)
	}

	filterParams := &twitter.StreamFilterParams{
		Track:         t.Keywords,
		StallWarnings: twitter.Bool(true),
	}
	log.Println("Creating the stream with filter params", t.Keywords)

	stream, err := client.Streams.Filter(filterParams)
	if err != nil {
		return nil, err
	}

	go demux.HandleChan(stream.Messages)

	return stream, nil
}
