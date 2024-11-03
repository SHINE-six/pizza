package services

import (
	"fmt"
	"log"
	"text/template"
	"bytes"
	"os"

	"Notification_Service/config"

	"gopkg.in/gomail.v2"
)

type VerificationEmailData struct {
	Name string
	VerificationLink string
}


func SendUserVerificationEmail(email string, username string, verificationToken string) error {
	cfg := config.LoadEmailConfig()

	// Generate the verification link
	verificationLink := fmt.Sprintf("%s?email=%s&token=%s", cfg.VerificationLinkPrefix, email, verificationToken)

	// Prepare the email body
	body, err := prepareEmailBody(username, verificationLink)
	if err != nil {
		log.Println(err)
		return err
	}

	err = sendEmail(email, body.String(), cfg.CompanyEmail, cfg.CompanyEmailPassword)
	if err != nil {
		log.Printf("Failed to send the verification email: %s", err)
		return err
	}

	return nil
}

func prepareEmailBody(username, verificationLink string) (*bytes.Buffer, error) {
	// list the directory now
	log.Printf("The directory is: %s", os.Getenv("PWD"))

	var body bytes.Buffer
	t, err := template.ParseFiles("./internal/pkg/verification_email_template.html")
	if err != nil {
		log.Printf("Failed to parse the email template: %s", err)
		return nil, err
	}

	t.Execute(&body, VerificationEmailData{Name: username, VerificationLink: verificationLink})

	return &body, nil
}

func sendEmail(recipientEmail, body string, companyEmail, companyEmailPassword string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", companyEmail)
	m.SetHeader("To", recipientEmail)
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", body)

	d := gomail.NewDialer("smtp.gmail.com", 587, companyEmail, companyEmailPassword)

	return d.DialAndSend(m)
}