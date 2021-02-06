package structs

// TwitterCredentials twitter credentials configuration
type TwitterCredentials struct {
	APIKey       string
	APISecret    string
	AccessToken  string
	AccessSecret string
}

// TwitterConfig represents Twitter configuration
type TwitterConfig struct {
	Credentials TwitterCredentials
	Keywords    []string
}

// KafkaConfig configuration for kafka
type KafkaConfig struct {
	BootstrapBrokers []string
	OutputTopic      string
}
