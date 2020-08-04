package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// ConnectDatabase creates a database connection for CRUD
func ConnectDatabase() *gorm.DB {
	const (
		POSTGRES_USER     = "golang"
		POSTGRES_PASSWORD = "googleGo"
		DB_NAME           = "todoGo"
		POSTGRES_HOST     = "localhost"
		POSTGRES_PORT     = 5432
	)

	postgres_conname := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		POSTGRES_HOST, POSTGRES_PORT, POSTGRES_USER, POSTGRES_PASSWORD, DB_NAME)

	db, err := gorm.Open("postgres", postgres_conname)
	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to database")
	}
	fmt.Println("Connected")

	db.AutoMigrate(&Task{})
	return db
}
