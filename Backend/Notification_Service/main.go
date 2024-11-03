package main

import (
	"encoding/json"
	"log"
	"Notification_Service/config"
	"Notification_Service/internal/services"

	"github.com/confluentinc/confluent-kafka-go/kafka"
    db_struct "pizza/db/struct"
)

func main() {
	cfg := config.LoadKafkaConfig()
    c, err := kafka.NewConsumer(&kafka.ConfigMap{
        "bootstrap.servers": cfg.BootstrapServers,
        "sasl.username":     cfg.SaslUsername,
        "sasl.password":    cfg.SaslPassword,
        "group.id": 		cfg.GroupID,
        "security.protocol": cfg.SecurityProtocol,
        "sasl.mechanisms":   cfg.SaslMechanisms,
        "auto.offset.reset": cfg.AutoOffsetReset,
	})

    if err != nil {
        log.Fatalf("Failed to create consumer: %s", err)
    }

    topic :=[]string{"user_registered", "order_created", "order_status_updated"}
    err = c.SubscribeTopics(topic, nil)
    if err != nil {
        log.Fatalf("Error subscribing to topic: %v", err)
    }
    defer c.Close()

    // Process messages
    for {
        msg, err := c.ReadMessage(-1)  // -1 means block until a message is received
        if err != nil {
            // Errors are informational and automatically handled by the consumer, so it is okay to continue
            log.Printf("Consumer error: %v (%v)\n", err, msg)
            continue
        }
        switch *msg.TopicPartition.Topic {
        case topic[0]:  // user_registered
            go handleUserRegistered(msg.Key, msg.Value)
        case topic[1]:  // order_created
            go handleOrderCreated(msg.Key, msg.Value)
        case topic[2]:  // order_status_updated
            go handleOrderStatusUpdated(msg.Key, msg.Value)
        default:
            log.Printf("Received message from unknown topic: %s\n", *msg.TopicPartition.Topic)
        }
    }
}

type UserRegisteredEvent struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	VerificationToken string `json:"verification_token"`
	Purpose  string `json:"purpose"`
}

func handleUserRegistered(key []byte, value []byte) {
    log.Printf("Handling topic: %s\n Message: %s\n", string(key), string(value))
	
	// Deserialize the event
	var event UserRegisteredEvent
	err := json.Unmarshal(value, &event)
	if err != nil {
		log.Fatalf("Failed to deserialize the event: %s", err)
	}

	// Send the verification email
	err = services.SendUserVerificationEmail(event.Email, event.Username, event.VerificationToken)
	if err != nil {
		log.Printf("Failed to send the verification email: %s", err)
	}
}


type OrderCreatedEvent struct {
    Email    string `json:"email"`
    Username string `json:"username"`
    Pizza   db_struct.Pizza `json:"pizza"`
    Purpose  string `json:"purpose"`
}

func handleOrderCreated(key []byte, value []byte) {
    log.Printf("Handling topic: %s\n Message: %s\n", string(key), string(value))
    
    // Deserialize the event
    var event OrderCreatedEvent
    err := json.Unmarshal(value, &event)
    if err != nil {
        log.Fatalf("Failed to deserialize the event: %s", err)
    }
    
    // Send the order created notification email
    err = services.SendOrderCreatedNotificationEmail(event.Email, event.Username, &event.Pizza)
    if err != nil {
        log.Printf("Failed to send the order created notification email: %s", err)
    }
}

type OrderStatusUpdatedEvent struct {
    Email           string `json:"email"`
    CustomerUsername string `json:"customerUsername"`
    DeliveryStaffID string `json:"deliveryStaffID"`
    OrderID         string `json:"orderId"`
    Status          string `json:"status"`
    Purpose         string `json:"purpose"`
}

func handleOrderStatusUpdated(key []byte, value []byte) {
    log.Printf("Handling topic: %s\n Message: %s\n", string(key), string(value))
    
    // Deserialize the event
    var event OrderStatusUpdatedEvent
    err := json.Unmarshal(value, &event)
    if err != nil {
        log.Fatalf("Failed to deserialize the event: %s", err)
    }
    
    // Send the order status updated notification email
    err = services.SendOrderStatusUpdatedNotificationEmail(event.Email, event.CustomerUsername, event.DeliveryStaffID, event.OrderID, event.Status)
    if err != nil {
        log.Printf("Failed to send the order status updated notification email: %s", err)
    }
}
