package config

import (
	"github.com/spf13/viper"
)


type KafkaConfig struct {
	BootstrapServers string		// Kafka broker address
	SaslUsername string		// SASL username
	SaslPassword string		// SASL password
	GroupID string
	SecurityProtocol string			// Security protocol: SASL_SSL, SSL, PLAINTEXT
	SaslMechanisms string				// SASL mechanism: PLAIN, SCRAM-SHA-256, SCRAM-SHA-512
	Acks string							// Acknowledgement: 0, 1, all; is used to control the durability of the messages
	AutoOffsetReset string		// Offset reset policy: earliest, latest, none: meaning if the consumer does not have an offset, it will start consuming from the earliest or latest message
}

func LoadKafkaConfig() *KafkaConfig {
	// viper.SetConfigFile("./config/kafka.env")
	// viper.AutomaticEnv()
	// err := viper.ReadInConfig()

	// if err != nil {
	// 	panic(err)
	// }

	viper.AutomaticEnv()

	return &KafkaConfig{
		BootstrapServers: viper.GetString("BOOTSTRAP_SERVERS"),
		SaslUsername: viper.GetString("SASL_USERNAME"),
		SaslPassword: viper.GetString("SASL_PASSWORD"),
		GroupID: viper.GetString("GROUP_ID"),
		SecurityProtocol: viper.GetString("SECURITY_PROTOCOL"),
		SaslMechanisms: viper.GetString("SASL_MECHANISMS"),
		Acks: viper.GetString("ACKS"),
		AutoOffsetReset: viper.GetString("AUTO_OFFSET_RESET"),
	}
}