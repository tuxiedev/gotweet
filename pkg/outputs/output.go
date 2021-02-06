package outputs

import (
	"github.com/dghubble/go-twitter/twitter"
)

// Output defines an output
type Output interface {
	Init() (error)
	Start(chan *twitter.Tweet)
	Stop()
}


// InitializeOutputs initializes and starts the output worker
func InitializeOutput(outputName string, outputConfig interface{}, 
	tweets chan *twitter.Tweet) (Output, error) {

	var initializedOutput Output
	switch outputName {
	case "console":
		initializedOutput = &Console{}
	case "kafka":
		initializedOutput = &Kafka{Config: outputConfig}
	}
	err := initializedOutput.Init()
	if err != nil {
		return nil, err
	}
	go initializedOutput.Start(tweets)
	return initializedOutput, nil
}