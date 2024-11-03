package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"errors"
	"log"
	"strconv"
	"gorm.io/gorm"
    "Order_Service/graph/model"

	db_struct "pizza/db/struct"
)

type Resolver struct{
	DB *gorm.DB
}

// Check if the order exists, if not create a new order
func checkOrderExists(db *gorm.DB, customerID string) (uint, error) {
    var order db_struct.Order

    customerIDUint, err := strconv.ParseUint(customerID, 10, 64)
    if err != nil { 
        log.Println("Error in checkOrderExists: ", err)
        return 0, err
    }

    result := db.Where("customer_id = ? AND status = ?", customerIDUint, "pending").First(&order)
    if errors.Is(result.Error, gorm.ErrRecordNotFound) {
        order = db_struct.Order{
            DeliveryStaffID: nil,   // Initially set to nil
            CustomerID: uint(customerIDUint),
            TotalPrice: 0,
            Status: "pending",
        }
        // Handle record creation with status pending
        result = db.Create(&order)
        if result.Error != nil {
            log.Println("Error creating order: ", result.Error)
            return 0, result.Error
        }
    } else if result.Error != nil {
        // Handle other potential errors from the initial db.First call
        log.Println("Error checking order exists: ", result.Error)
        return 0, result.Error
    }
    return order.ID, nil
}

type PricedItem interface {
    GetPrice() float64
}


func dbToppingsPrice(dbToppings []db_struct.Topping) float64 {
    var totalPrice float64
    for _, topping := range dbToppings {
        totalPrice += topping.Price
    }

    return totalPrice
}

func toppingsPrice(toppings []model.Topping) float64 {
    var totalPrice float64
    for _, topping := range toppings {
        totalPrice += topping.Price
    }

    return totalPrice
}