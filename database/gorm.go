package database

import (
	"fmt"
	"log"
	"tugas_2/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	user     = "postgres"
	password = "Zekkenyope121*"
	dbPort   = 5000
	dbName   = "assignment_2"
)

var db *gorm.DB

func GormInit() {

	config := fmt.Sprintf("host=%s user=%s password=%s port=%d dbname=%s sslmode=disable", host, user, password, dbPort, dbName)

	// Connect to the database
	var err error
	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Migrate the database schema
	db.AutoMigrate(&models.Order{}, &models.Item{})
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return db
}
