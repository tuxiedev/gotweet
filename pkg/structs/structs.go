package structs

// TwitterCredentials twitter credentials configuration
type TwitterCredentials struct {
	APIKey         string
	APISecret      string
	ConsumerKey    string
	ConsumerSecret string
}

// KafkaConfig configuration for kafka
type KafkaConfig struct {
	BootstrapBrokers string
	OutputTopic      string
}
