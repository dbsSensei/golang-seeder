package main

import (
	"fmt"
	"github.com/dbsSensei/golang-seeder/models"
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

// Seeder function
func seedUsers(db *gorm.DB) {
	users := []models.User{
		{
			Name:      "John Doe",
			Email:     "john.doe@example.com",
			Age:       30,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Name:      "Jane Doe",
			Email:     "jane.doe@example.com",
			Age:       25,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	for _, user := range users {
		// generate and embed uuid to user
		id := uuid.New()
		user.ID = id

		db.Create(&user)
		fmt.Println("user", user.ID)
	}
}

// exec di main function
func main() {
	dsn := "root:password@tcp(127.0.0.1:3306)/golang_seeder?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	if err := db.Debug().AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	seedUsers(db)
}
