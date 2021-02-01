package console

import (
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
)

// Output Runs the console output
func Output(tweetChannel chan *twitter.Tweet) {
	for tweet := range tweetChannel {
		fmt.Println(tweet.Text)
	}
}
