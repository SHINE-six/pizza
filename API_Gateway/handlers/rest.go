package handlers

import (
	"API_Gateway/microservices"

	"github.com/gin-gonic/gin"
)

func HandleRest(r *gin.RouterGroup) {
	// User Service
	r.POST("/v1/user/login", microservices.AuthUser)
	r.POST("/v1/user/signup", microservices.CreateUser)
	r.GET("/v1/user/verify_email", microservices.VerifyUser)
	r.GET("/v1/staff/verify_email", microservices.VerifyStaff)
	// r.GET("/v1/verifyCookie", microservices.VerifyCookie)

	// Menu Service
	// r.GET("/menu", microservices.GetMenu) in graphql

	// Order Service
	// r.GET("/order", microservices.GetOrder)
	// r.POST("/order", microservices.CreateOrder)
	// r.PUT("/order", microservices.UpdateOrder)

	// Delivery Service
	// r.GET("/delivery", microservices.GetDelivery)
}
