package outputs

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/Shopify/sarama"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/tuxiedev/gotweet/pkg/structs"
)

func TestKafka(t *testing.T) {
	t.Run("Initialize a kafka output", func(t *testing.T) {
		tweets := make(chan *twitter.Tweet)

		clusterAdmin, err := sarama.NewClusterAdmin([]string{"localhost:9092"}, nil)

		clusterAdmin.CreateTopic("tweets", &sarama.TopicDetail{NumPartitions: 1, ReplicationFactor: 1}, false)

		output := &Kafka{
			Config: structs.KafkaConfig{
				BootstrapBrokers: []string{"localhost:9092"},
				OutputTopic:      "tweets",
			},
		}
		output.Init()
		go output.Start(tweets)

		log.Println("Started output")

		consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
		if err != nil {
			panic(err)
		}
		log.Println("Initialized consumer")

		defer func() {
			if err := consumer.Close(); err != nil {
				t.Fail()
			}
		}()

		partitionConsumer, err := consumer.ConsumePartition("tweets", 0, sarama.OffsetNewest)
		log.Println("Initialized partitionConsumer")
		if err != nil {
			panic(err)
		}

		defer func() {
			if err := partitionConsumer.Close(); err != nil {
				t.FailNow()
			}
		}()

		doneCh := make(chan bool)
		var receivedMessage *sarama.ConsumerMessage

		go func() {
			for {
				select {
				case msg := <-partitionConsumer.Messages():
					receivedMessage = msg
					doneCh <- true
				}
			}
		}()

		tweetToVerify := &twitter.Tweet{
			Text: "this is a test tweet",
		}

		tweets <- tweetToVerify

		log.Println("Waiting to consume message")
		<-doneCh
		log.Printf("received message %v\n", string(receivedMessage.Value))
		var receivedTweet twitter.Tweet

		err = json.Unmarshal(receivedMessage.Value, &receivedTweet)

		if err != nil {
			log.Fatalf("Could not unmarshal messsage into tweet %v\n", err)
			t.Fail()
		}
		if receivedTweet.Text != tweetToVerify.Text {
			log.Fatalf("%v != %v \n", receivedTweet.Text, tweetToVerify.Text)
		}

		output.Stop()
		clusterAdmin.DeleteTopic("tweets")
		clusterAdmin.Close()
	})
}
