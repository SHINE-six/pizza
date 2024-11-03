package services

import (
	"bytes"
	"log"
	"strconv"
	"strings"
	"text/template"

	"Notification_Service/config"

	db_struct "pizza/db/struct"

	"gopkg.in/gomail.v2"
)

type OrderCreatedEmailData struct {
	Name string
	OrderID string
	OrderDate string
	PizzaBase string
	PizzaBasePrice float64
	PizzaToppings_Name_Price string
	PizzaSize string
	PizzaSizeMultiplier float64
	TotalPrice float64
}


func SendOrderCreatedNotificationEmail(email string, username string, pizza *db_struct.Pizza) error {
	cfg := config.LoadEmailConfig()

	// Prepare the email body
	body, err := prepareOrderCreatedEmailBody(username, pizza)
	if err != nil {
		log.Println(err)
		return err
	}

	err = sendOrderCreatedEmail(email, body.String(), cfg.CompanyEmail, cfg.CompanyEmailPassword)
	if err != nil {
		log.Printf("Failed to send the verification email: %s", err)
		return err
	}

	return nil
}

func prepareOrderCreatedEmailBody(username string, pizza *db_struct.Pizza) (*bytes.Buffer, error) {
	var body bytes.Buffer
	t, err := template.ParseFiles("./internal/pkg/order_created_noti_template.html")
	if err != nil {
		log.Printf("Failed to parse the email template: %s", err)
		return nil, err
	}

	orderID := strconv.Itoa(int(pizza.ID))
	orderDate:= pizza.CreatedAt.String()
	baseName := pizza.Base.Name
	basePrice := pizza.Base.Price
	sizeName := pizza.Size.Name
	sizeMultiplier := pizza.Size.Multiplier
	// break down the toppings into a slice of strings
	toppings_name_price := make([]string, len(pizza.Toppings))
	for i, topping := range pizza.Toppings {
		toppings_name_price[i] = topping.Name + "    RM " + strconv.FormatFloat(topping.Price, 'f', 2, 64)
	}
	joinedToppings := strings.Join(toppings_name_price, "<br>")
	price := pizza.Price


	t.Execute(&body, OrderCreatedEmailData{
		Name: username,
		OrderID: orderID,
		OrderDate: orderDate,
		PizzaBase: baseName,
		PizzaBasePrice: basePrice,
		PizzaToppings_Name_Price: joinedToppings,
		PizzaSize: sizeName,
		PizzaSizeMultiplier: sizeMultiplier,
		TotalPrice: price,
	})

	return &body, nil
}

func sendOrderCreatedEmail(recipientEmail, body string, companyEmail, companyEmailPassword string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", companyEmail)
	m.SetHeader("To", recipientEmail)
	m.SetHeader("Subject", "Order Created")
	m.SetBody("text/html", body)

	d := gomail.NewDialer("smtp.gmail.com", 587, companyEmail, companyEmailPassword)

	return d.DialAndSend(m)
}