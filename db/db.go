package db

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"
)

var DB *gorm.DB

func Connect() {

	/*dsn := "host=localhost user=stylianoskoulinas password=zPMa]W?dxY5M dbname=musicdb port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})*/

	db, err := gorm.Open(sqlite.Open("records.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	DB = db
	fmt.Println("Database connected")
}
