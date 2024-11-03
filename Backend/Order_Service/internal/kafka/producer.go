package kafka

import (
	"Order_Service/config"
	"Order_Service/internal/jwt"
	"context"
	"encoding/json"
	"log"

	db_struct "pizza/db/struct"
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

func ProduceOrderCreatedEvent(ctx context.Context, email string, username string, pizza *db_struct.Pizza) (int, string) {
	topic := "order_created"

	if email == "" || username == "" {
		resp, err := jwt.DecodeJWT(ctx, "")
		if resp["email"] == nil || resp["username"] == nil || err != nil {
			log.Printf("Failed to extract the email, username, or verification token from the token")
			return 1, "Failed to extract the email, username, or verification token from the token"
		}
		email = resp["email"].(string)
		username = resp["username"].(string)

		log.Printf("Email: %s, Username: %s\n", email, username)
	}

	// Define the UserRegisteredEvent struct
	type OrderCreatedEvent struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Pizza    db_struct.Pizza `json:"pizza"`
		Purpose  string `json:"purpose"`
	}

	// Create a new instance of the UserRegisteredEvent struct
	event := OrderCreatedEvent{
		Email:    email,
		Username: username,
		Pizza:    *pizza,
		Purpose:  "Send order created notification email",   // TODO: Fix the number of available purposes in the future from a custom package, e.g., kafka.purpose.send_verification_email, kafka.purpose.send_password_reset_email
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
		log.Printf("Produced event to topic %s: msg = %s\n", topic, string(serializedEvent))
		return 0, ""
	}
}

func ProduceOrderDeliveringEvent(email string, customerUsername string, deliveryStaffID string, orderId string) {
	ProduceOrderStatusUpdatedEvent(email, customerUsername, deliveryStaffID, orderId, "delivering", "Send order delivering notification email")
}

func ProduceOrderStatusUpdatedEvent(email string, customerUsername string, deliveryStaffID string, orderId string, status string, Purpose string) {
	topic := "order_status_updated"

	// Define the OrderStatusUpdatedEvent struct
	type OrderStatusUpdatedEvent struct {
		Email           string `json:"email"`
		CustomerUsername string `json:"customerUsername"`
		DeliveryStaffID string `json:"deliveryStaffID"`
		OrderID         string `json:"orderId"`
		Status          string `json:"status"`
		Purpose         string `json:"purpose"`
	}

	// Create a new instance of the OrderStatusUpdatedEvent struct
	event := OrderStatusUpdatedEvent{
		Email:           email,
		CustomerUsername: customerUsername,
		DeliveryStaffID: deliveryStaffID,
		OrderID:         orderId,
		Status:          status,
		Purpose:         Purpose,
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
	} else {
		log.Printf("Produced event to topic %s: msg = %s\n", topic, string(serializedEvent))
	}
}