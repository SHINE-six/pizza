package main

import (
	"API_Gateway/config"
	"API_Gateway/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := setupRouter()
	cfg := config.LoadConfig()
	println(cfg.ServerPort)

	registerEndpoints(r)
	
	r.Run(cfg.ServerPort)
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(CORSMiddleware())

	return r
}

func registerEndpoints(r *gin.Engine) {
	// REST API routes
    handlers.HandleRest(r.Group("/rest"))

    // GraphQL route
    r.POST("/graphql", handlers.HandleGraphql)
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