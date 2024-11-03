package jwt

import (
	"fmt"
	"log"
	"context"
	"Order_Service/config"

	"github.com/golang-jwt/jwt/v5"
)

func getToken(ctx context.Context) string {
	token, ok := ctx.Value(TokenKey).(string)
	if !ok {
		log.Printf("Token not found")
		return ""
	}
	return token
}
// DecodeJWT decodes the JWT token to it's json nature value
func DecodeJWT(ctx context.Context, tokenString string) (jwt.MapClaims, error) {
	if tokenString == "" {
		tokenString = getToken(ctx)
	}
	// Parse the JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Get the secret key from the environment
		secretKey := config.LoadJWTConfig().SecretKey
		if secretKey == "" {
			log.Fatalf("JWT_SECRET environment variable is not set")
		}

		// Return the secret key
		return []byte(secretKey), nil
	})

	// Check if the token is valid
	if err != nil {
		return nil, fmt.Errorf("failed to parse the token: %s", err)
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, fmt.Errorf("token is not valid")
	}

	// Get the claims from the token
	claims, ok := token.Claims.(jwt.MapClaims)
	log.Printf("Claims: %v", claims)
	if !ok {
		return nil, fmt.Errorf("failed to get claims from the token")
	}

	// Return the claims without the map array
	return claims, nil
}