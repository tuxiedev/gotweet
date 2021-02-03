package outputs

import (
	"github.com/dghubble/go-twitter/twitter"
)

// Output defines an output
type Output interface {
	Init()
	Start(chan *twitter.Tweet)
	Stop()
}
