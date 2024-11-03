package config

import (
	"log"

    "github.com/spf13/viper"
)

type Config struct {
    ServerPort   string
    UserServiceURL string
    Service2Addr string
}

// LoadConfig loads the configuration from the .env file
func LoadConfig() *Config {
	// viper.SetConfigFile("./config/.env")
	// viper.AutomaticEnv()
	// err := viper.ReadInConfig()

	// if err != nil {
	// 	panic(err)
	// }

	// Set up viper to read environment variables
	viper.AutomaticEnv()

	// Set default values (optional)
	viper.SetDefault("SERVER_PORT", ":8080")
	viper.SetDefault("USER_SERVICE_URL", "user_service-backend:50051")

	// Read environment variables
	serverPort := viper.GetString("SERVER_PORT")
	userServiceURL := viper.GetString("USER_SERVICE_URL")

	// Log the loaded configurations (for debugging purposes)
	log.Printf("Loaded Configurations: SERVER_PORT=%s, USER_SERVICE_URL=%s",
		serverPort, userServiceURL)


	return &Config{
		ServerPort:   serverPort,
		UserServiceURL:   userServiceURL,
	}
}
