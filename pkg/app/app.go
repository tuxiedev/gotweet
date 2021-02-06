package app

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	dgt "github.com/dghubble/go-twitter/twitter"
	"github.com/tuxiedev/gotweet/pkg/outputs"
	"github.com/tuxiedev/gotweet/pkg/structs"
	"github.com/tuxiedev/gotweet/pkg/twitter"
)

// RunApp starts the main loop of the app
func RunApp(t structs.TwitterConfig, outputName string, outputConfig interface{}) {

	tweets := make(chan *dgt.Tweet)

	stream, err := twitter.StartConsumingFromTwitter(t, tweets)

	if err != nil {
		log.Fatalf("Failed to initialize twitter stream %v\n", err)
	}

	output, err := outputs.InitializeOutput(outputName, outputConfig, tweets)

	if err != nil {
		log.Fatalf("Failed to initialize output %\n", err)
	}

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)

	fmt.Println("Stopping Stream...")
	output.Stop()
	stream.Stop()
	close(tweets)

}
