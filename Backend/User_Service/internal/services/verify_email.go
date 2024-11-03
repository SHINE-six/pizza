package services

import (
	"crypto/rand"
	"encoding/base64"
)


func SendVerificationEmail(email string, username string, password string) (status int, message string) {
	// Generate a verification token
    verificationToken, _ := generateVerificationToken()

	// Add the email and verification token to the database table - verification_tokens
	// The table should have the following columns: email, token, created_at, updated_at
	statusCode, err := addVerificationTokenToDB(email, verificationToken, username, password)
	if statusCode != 0 {
		return statusCode, err.Error()
	}
	
	// Produce a UserRegisteredEvent to the Kafka topic
	statusCode, msg := ProduceUserRegisteredEvent(email, username, verificationToken)
	if statusCode != 0 {
		return 1, msg
	}

    return 0, "Verification email sent"
}

func generateVerificationToken() (string, error) {
	// Generate a random 16 digit string
	token := make([]byte, 16)
	_, err := rand.Read(token)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(token), nil
}