package db

import (
	"log"
	"Order_Service/config"

	"gorm.io/gorm"
	"gorm.io/driver/postgres"
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