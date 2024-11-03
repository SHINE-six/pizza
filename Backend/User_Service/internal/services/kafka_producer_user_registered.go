package services

import (
	"log"
	"encoding/json"
	"User_Service/config"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// Global variable
var p *kafka.Producer

func init() {
	cfg := config.LoadKafkaConfig()

	var err error
	p, err = kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": cfg.BootstrapServers,
		"sasl.username":     cfg.SaslUsername,
		"sasl.password":     cfg.SaslPassword,
		"security.protocol": cfg.SecurityProtocol,
		"sasl.mechanisms":   cfg.SaslMechanisms,
		"acks":              cfg.Acks,	
	})

	if err != nil {
		log.Fatalf("Failed to create producer: %s", err)
	}
	log.Printf("Kafka producer created successfully\n")
}

func ProduceUserRegisteredEvent(email string, username string, verificationToken string) (int, string) {
	topic := "user_registered"

	// Define the UserRegisteredEvent struct
	type UserRegisteredEvent struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		VerificationToken string `json:"verification_token"`
		Purpose  string `json:"purpose"`
	}

	// Create a new instance of the UserRegisteredEvent struct
	event := UserRegisteredEvent{
		Email:    email,
		Username: username,
		VerificationToken: verificationToken,
		Purpose:  "Send user verification email",   // TODO: Fix the number of available purposes in the future from a custom package, e.g., kafka.purpose.send_verification_email, kafka.purpose.send_password_reset_email
	}

	// Serialize the event
	serializedEvent, err := json.Marshal(event)
	if err != nil {
		log.Fatalf("Failed to serialize the event: %s", err)
	}

	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            []byte(email),
		Value:          serializedEvent,
	}, nil)

	if err != nil {
		log.Printf("Failed to produce message: %s\n", err)
		return 2, "Failed to produce message"
	} else {
		log.Printf("Produced event to topic %s: key = %-10s value = %s\n", topic, email, username)
		return 0, ""
	}
}