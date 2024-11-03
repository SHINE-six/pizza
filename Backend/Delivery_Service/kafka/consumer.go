package kafka

import (
	"Delivery_Service/config"
	"encoding/json"
	// "encoding/json"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// Global variable
var c *kafka.Consumer

func init() {
	cfg := config.LoadKafkaConfig()

	var err error
	c, err = kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": cfg.BootstrapServers,
		"sasl.username":     cfg.SaslUsername,
		"sasl.password":     cfg.SaslPassword,
		"group.id":          cfg.GroupID,
		"security.protocol": cfg.SecurityProtocol,
		"sasl.mechanisms":   cfg.SaslMechanisms,
		"auto.offset.reset": cfg.AutoOffsetReset,
	})

	if err != nil {
		log.Fatalf("Failed to create consumer: %s", err)
	}
	log.Printf("Kafka consumer created successfully\n")
}


func ConsumeLocationEvent(deliveryStaffID string, callback func([]byte)) {
	topic := "location_data"

	err := c.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		log.Fatalf("Error subscribing to topic: %v", err)
	}
	defer c.Close()

	// Process messages
	for {
		msg, err := c.ReadMessage(-1) // -1 means block until a message is received
		if err != nil {
			// Errors are informational and automatically handled by the consumer, so it is okay to continue
			log.Printf("Consumer error: %v (%v)\n", err, msg)
			continue
		}

		var locationEvent LocationUpdateEvent
		if err = json.Unmarshal(msg.Value, &locationEvent); err != nil {
			log.Printf("Failed to unmarshal the message: %s\n", err)
			continue // Skip the rest of the loop
		}

		if locationEvent.StaffId == deliveryStaffID {
			log.Printf("Received message: %s\n", string(msg.Value))
			callback(msg.Value)
		}
	}
}