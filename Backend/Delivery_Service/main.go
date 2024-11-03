package main

import (
	"Delivery_Service/handlers"

	"github.com/gin-gonic/gin"
)

// Websocket Server
func main() {
	r := setupRouter()
	
	r.GET("/ws/delivery", handlers.HandleWebsocketDelivery)
	r.GET("/ws/customer", handlers.HandleWebsocketCustomer)

	r.Run(":8081")
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(CORSMiddleware())

	return r
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}