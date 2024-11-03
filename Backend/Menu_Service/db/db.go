package db

import (
	"log"
	"Menu_Service/config"

	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	db_struct "pizza/db/struct"
)

var db *gorm.DB

func Connect() {
	var err error
	cfg := config.LoadDBConfig()
	dsn := cfg.PostgresDatabaseURL
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	log.Println("Connected to the database")

	AutoMigrate()   //* Only for the first time
}

func AutoMigrate() {
	log.Println("Auto migrating the database")
	if err := db.AutoMigrate(&db_struct.Base{}, &db_struct.Size{}, &db_struct.Topping{}, &db_struct.Customer{}, &db_struct.DeliveryStaff{}, &db_struct.Order{}, &db_struct.Pizza{}); err != nil {
		log.Fatalf("Error automigrating the database: %v", err)
	}
	log.Println("Auto migration completed")
}

func Close() {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Error getting sql.DB from GORM: %v", err)
	}
	if err := sqlDB.Close(); err != nil {
		log.Fatalf("Error closing the database connection: %v", err)
	}
	log.Println("Closed the database connection")
}


func GetDB() *gorm.DB {
	return db
}