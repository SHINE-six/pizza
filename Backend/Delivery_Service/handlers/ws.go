package handlers

import (
	"Delivery_Service/kafka"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

type ReceivedMessage struct {
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// HandleWebsocket handles websocket connections
func HandleWebsocketDelivery(c *gin.Context) {
	staffId := c.Query("staffId")
	log.Println("Websocket connection established" + staffId)

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Error upgrading to websocket connection: ", err)
		return
	}
	defer conn.Close()

	log.Println("Websocket connection established" + staffId)

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message: ", err)
			break
		}

		// Attempt to unmarshal the message into the ReceivedMessage struct
		var receivedMessage ReceivedMessage
		err = json.Unmarshal(msg, &receivedMessage)
		if err != nil {
			log.Printf("Failed to unmarshal the message: %s\n", err)
			continue // Skip the rest of the loop
		}

		if receivedMessage.Latitude == 0.0 && receivedMessage.Longitude == 0.0 {
			log.Printf("Invalid latitude and longitude: %f, %f\n", receivedMessage.Latitude, receivedMessage.Longitude)
		} else {
			// Produce the message to Kafka
			kafka.ProduceLocationEvent(c, staffId, receivedMessage.Latitude, receivedMessage.Longitude)
		}


	}
}

func HandleWebsocketCustomer(c *gin.Context) {
	deliveryStaffID := c.Query("deliveryStaffID")
	log.Println("Websocket connection established for customer " + deliveryStaffID)

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Error upgrading to websocket connection: ", err)
		return
	}
	defer conn.Close()

	log.Println("Websocket connection established" + deliveryStaffID)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	done := make(chan struct{})   // Channel to notify the main goroutine that the consumer has closed

	// Write message to the websocket connection with the kafka consumer
	// Consume the message from Kafka
	go func() {
		defer close(done)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				kafka.ConsumeLocationEvent(deliveryStaffID, func(msg []byte) {
					err := conn.WriteMessage(websocket.TextMessage, msg)
					if err != nil {
						log.Println("Error writing message: ", err)
						return
					}
				})
			}
		}
	}()

	for {
		if _, _, err := conn.ReadMessage(); err != nil {
			log.Println("Error reading message: ", err)
			break
		}
	}
}



	// // Write message to the websocket connection every 5 seconds
	// ticker := time.NewTicker(5 * time.Second)
	// defer ticker.Stop()

	// go func() {
	// 	for range ticker.C {
	// 		toReturnMsg := "Delivery staff " + deliveryStaffID + " is on the way" + time.Now().String()
	// 		err := conn.WriteMessage(websocket.TextMessage, []byte(toReturnMsg))
	// 		if err != nil {
	// 			log.Println("Error writing message: ", err)
	// 			return
	// 		}
	// 	}
	// }()