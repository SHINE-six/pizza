package kafka

import (
	"Delivery_Service/config"
	"context"
	"encoding/json"
	"log"

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

// Define the LocationUpdateEvent struct
type LocationUpdateEvent struct {
	StaffId   string `json:"staffId"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func ProduceLocationEvent(ctx context.Context, staffId string, latitude float64, longitude float64) {
	topic := "location_data"


	// Create a new instance of the LocationUpdateEvent struct
	event := LocationUpdateEvent{
		StaffId:   staffId,
		Latitude:  latitude,
		Longitude: longitude,
	}

	// Marshal the event
	eventBytes, err := json.Marshal(event)
	if err != nil {
		log.Printf("Failed to marshal the event: %s\n", err)
		return
	}

	// Produce the event
	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          eventBytes,
	}, nil)

	if err != nil {
		log.Printf("Failed to produce the event: %s\n", err)
	} else {
		log.Printf("Produced the event: %s\n", eventBytes)
	}
}