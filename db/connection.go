package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = "host=localhost port=5435 user=postgres dbname=postgres password=postgres"
var DB *gorm.DB


func DBConnection() {
	var error error
	DB, error = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if error != nil {
		// panic("Failed to connect to database!")
		log.Fatal("Failed to connect to database!", error)
	} else{
		log.Println("Connection to database established!")
	}

}
