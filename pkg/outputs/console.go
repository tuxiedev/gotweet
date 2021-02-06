package outputs

import (
	"encoding/json"
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
)

// Console outputs tweets json to console
type Console struct {
}

// Init initializes console output
func (c *Console) Init() error {
	fmt.Println("Starting console output")
	return nil
}

// Start starts the console output listening to a channel of tweets
func (c *Console) Start(tweetChannel chan *twitter.Tweet) {
	for tweet := range tweetChannel {
		tbytes, err := json.Marshal(tweet)
		if err != nil {
			fmt.Println("ERR: error decoding json from tweet")
		}
		fmt.Println(string(tbytes))
	}
}

// Stop prints a final message before exiting
func (c *Console) Stop() {
	fmt.Println("Closing console output")
}
