package graph

import (
	"gorm.io/gorm"

	// db_struct "pizza/db/struct"
)

type Resolver struct{
	DB *gorm.DB
}