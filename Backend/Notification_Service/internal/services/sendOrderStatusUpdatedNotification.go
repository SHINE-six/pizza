package services

import (
	"bytes"
	"log"
	"fmt"
	"text/template"
	"Notification_Service/config"
)

type OrderDeliveringEmailData struct {
	Name string
	OrderID string
	TrackingLink string
}

func SendOrderStatusUpdatedNotificationEmail(email string, customerUsername string, deliveryStaffID string, orderId string, status string) error {
	cfg := config.LoadEmailConfig()

	// Generate the tracking link
	trackingLink := fmt.Sprintf("%s?deliveryStaffID=%s", cfg.TrackingLinkPrefix, deliveryStaffID)

	// Prepare the email body
	var body *bytes.Buffer
	var err error
	
	if status == "delivering" {
		body, err = prepareOrderDeliveringEmailBody(customerUsername, orderId, trackingLink)
	} else {
		// Havent implemented the other status yet
	}

	if err != nil {
		log.Println(err)
		return err
	}

	err = sendEmail(email, body.String(), cfg.CompanyEmail, cfg.CompanyEmailPassword)
	if err != nil {
		log.Printf("Failed to send the order status updated notification email: %s", err)
		return err
	}

	return nil
}

func prepareOrderDeliveringEmailBody(customerUsername string, orderId string, trackingLink string) (*bytes.Buffer, error) {
	var body bytes.Buffer
	t, err := template.ParseFiles("./internal/pkg/order_delivering_email_template.html")
	if err != nil {
		log.Printf("Failed to parse the email template: %s", err)
		return nil, err
	}

	t.Execute(&body, OrderDeliveringEmailData{Name: customerUsername, OrderID: orderId, TrackingLink: trackingLink})

	return &body, nil
}