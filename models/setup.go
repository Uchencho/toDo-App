package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func SetupModels() *gorm.DB {
	const (
		POSTGRES_USER     = "golang"
		POSTGRES_PASSWORD = "googleGo"
		DB_NAME           = "todoGo"
		POSTGRES_HOST     = "localhost"
		POSTGRES_PORT     = 5432
	)

	postgres_conname := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		POSTGRES_HOST, POSTGRES_PORT, POSTGRES_USER, POSTGRES_PASSWORD, DB_NAME)

	fmt.Println("conname is\t\t", postgres_conname)

	db, err := gorm.Open("postgres", postgres_conname)
	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to database")
	}
	defer db.Close()

	db.AutoMigrate(&Task{})

	// Initialize record
	m := Task{
		Name:        "Uchenna",
		Description: "First record into postgres from Go",
		StartTime:   "03-08-2020",
		Alarm:       false,
	}
	db.Create(&m)
	fmt.Println("Successfully connected!")
	return db
}
