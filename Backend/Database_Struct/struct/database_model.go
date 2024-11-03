package model

import (
    "gorm.io/gorm"
)

// Dimension Table for display menu purpose
type Base struct {
    gorm.Model
	Name      string
	Price float64
}

// Dimension Table for display menu purpose
type Size struct {
	gorm.Model
	Name       string
	Multiplier float64
}

// Dimension Table for display menu purpose
type Topping struct {
	gorm.Model
	Name  string
	Price float64
}

// Fact Table for order purpose
type Pizza struct {
    gorm.Model
    OrderID   uint   `gorm:"index:idx_order_id"` // Pizza belongs to one order 
    BaseID    uint `gorm:"index:idx_base_id"`
    Base      Base  // Pizza belongs to one base
    SizeID    uint `gorm:"index:idx_size_id"`
    Size           Size  // Pizza belongs to one size
    Toppings       []Topping  `gorm:"many2many:pizza_toppings;"`// Has many Toppings, through pizza_toppings join table
    Price float64
}

// Fact Table for order purpose
type Order struct {
    gorm.Model
    DeliveryStaffID *uint `gorm:"index:idx_delivery_staff_id"`  // Order belongs to one delivery staff
    CustomerID uint  `gorm:"index:idx_customer_id"`
    Customer  Customer  // Order belongs to one customer
    Pizzas    []Pizza  // One order can have many pizzas or 1 pizza
    Status    string
    TotalPrice float64   //`gorm:"index:idx_total_price"` //TODO: but the index will only be implemented in future when necessary
}

// Fact Table for account purpose
type Customer struct {
    gorm.Model
    Email string `gorm:"index:idx_email,unique"`
    Username  string
    Password  string
}

// Fact Table for account purpose
type DeliveryStaff struct {
    gorm.Model
    Name string
    Email string `gorm:"index:idx_email,unique"`
    Orders []Order // One delivery staff can have many orders
}
    