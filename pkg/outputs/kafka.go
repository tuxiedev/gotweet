package outputs

import (
	"encoding/json"
	"log"

	"github.com/Shopify/sarama"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/tuxiedev/gotweet/pkg/structs"
)

// Kafka sends tweet to Kafka
type Kafka struct {
	Config   interface{}
	config   structs.KafkaConfig
	producer sarama.AsyncProducer
}

// Init initializes an async Kafka producer
func (k *Kafka) Init() {
	k.config = k.Config.(structs.KafkaConfig)
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal // Wait for all in-sync replicas to ack the message
	config.Producer.Retry.Max = 10                     // Retry up to 10 times to produce the message
	config.Producer.Return.Successes = true

	producer, err := sarama.NewAsyncProducer(k.config.BootstrapBrokers, config)
	if err != nil {
		log.Fatalln("Failed to start Sarama producer:", err)
	}

	k.producer = producer

	go func() {
		for range k.producer.Successes() {
			// noop
		}
	}()

	go func() {
		for err := range producer.Errors() {
			log.Println(err)
		}
	}()

}

func (k *Kafka) sendTweetToKafka(tweet []byte) {
	select {
	case k.producer.Input() <- &sarama.ProducerMessage{
		Topic: k.config.OutputTopic,
		Value: sarama.ByteEncoder(tweet),
	}:
	default:
	}
}

// Start starts the output to Kafka
func (k *Kafka) Start(tweetChannel chan *twitter.Tweet) {
	for {
		select {
		case tweet := <-tweetChannel:
			tbytes, err := json.Marshal(tweet)
			if err != nil {
				log.Println("ERR: error decoding json from tweet")
			}
			go k.sendTweetToKafka(tbytes)
		default:
		}
	}
}

// Stop closes the Kafka producer on interrupt
func (k *Kafka) Stop() {
	k.producer.Close()
}
