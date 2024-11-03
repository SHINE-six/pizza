package config

import (
    "github.com/spf13/viper"
)

type ServerConfig struct {
    ServerPort   string
    APIGatewayURL string
	NetworkType string
}

// LoadServerConfig loads the server configuration from the .env file
func LoadServerConfig() *ServerConfig {
	// viper.SetConfigFile("./config/server.env")
	// viper.AutomaticEnv()
	// err := viper.ReadInConfig()

	// if err != nil {
	// 	panic(err)
	// }

	// Set up viper to read environment variables
	viper.AutomaticEnv()

	// Set default values (optional)
	viper.SetDefault("SERVER_PORT", ":50051")
	viper.SetDefault("API_GATEWAY_URL", "localhost:8080")
	viper.SetDefault("NETWORK_TYPE", "tcp")

	// Read environment variables
	serverPort := viper.GetString("SERVER_PORT")
	apiGatewayURL := viper.GetString("API_GATEWAY_URL")
	networkType := viper.GetString("NETWORK_TYPE")

	return &ServerConfig{
		ServerPort:  serverPort,
		APIGatewayURL: apiGatewayURL,
		NetworkType: networkType,
	}
}

type JWTConfig struct {
	SecretKey string
}

func LoadJWTConfig() *JWTConfig {
	// viper.SetConfigFile("./config/jwt.env")
	// viper.AutomaticEnv()
	// err := viper.ReadInConfig()

	// if err != nil {
	// 	panic(err)
	// }

	viper.AutomaticEnv()

	secretKey := viper.GetString("SECRETKEY")

	return &JWTConfig{
		SecretKey: secretKey,
	}
}

type KafkaConfig struct {
	BootstrapServers string		// Kafka broker address
	SaslUsername string		// SASL username
	SaslPassword string		// SASL password
	SecurityProtocol string			// Security protocol: SASL_SSL, SSL, PLAINTEXT
	SaslMechanisms string				// SASL mechanism: PLAIN, SCRAM-SHA-256, SCRAM-SHA-512
	Acks string							// Acknowledgement: 0, 1, all; is used to control the durability of the messages
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
		SecurityProtocol: viper.GetString("SECURITY_PROTOCOL"),
		SaslMechanisms: viper.GetString("SASL_MECHANISMS"),
		Acks: viper.GetString("ACKS"),
	}
}
type dbConfig struct {
	PostgresDatabaseURL string
}

func LoadDBConfig() *dbConfig {
	// viper.SetConfigFile("./config/db_env.json")
	// viper.AutomaticEnv()
	// err := viper.ReadInConfig()

	// if err != nil {
	// 	panic(err)
	// }

	viper.AutomaticEnv()

	return &dbConfig{
		PostgresDatabaseURL: viper.GetString("POSTGRES_DATABASE_URL"),
	}
}
