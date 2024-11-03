package services

import (
	"fmt"
	"time"
	"User_Service/config"

	"github.com/golang-jwt/jwt/v5"
)

func CreateJWTToken(email string) (string, error) {
	fmt.Println("Creating a JWT token for the user")
	
	cfg := config.LoadJWTConfig()
	secretKey := []byte(cfg.SecretKey)
	// Create a JWT token for the user
	// The token should contain the email and username
	username := getUsernameFromEmail(email)
	// The token should expire in 30 minutes
	claims := jwt.MapClaims{
		"customerId": getCustomerIDFromEmail(email),
		"email": email,
		"username": username,
		"exp": time.Now().Add(time.Minute * 30).Unix(),  // 30 minutes
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println("Error signing the JWT token: ", err)
		return "", err
	}

	return tokenString, err
}