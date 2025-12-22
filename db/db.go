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

<<<<<<< HEAD
	/*dsn := "host=localhost user=stylianoskoulinas password=zPMa]W?dxY5M dbname=musicdb port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})*/
=======
	//PostgreSQL connection string
	/*dsn := "host=localhost user=postgres password= dbname=postgres port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}) */
>>>>>>> e5eb7b459f4c676afcdc753128f7862b93da8e41

	db, err := gorm.Open(sqlite.Open("records.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	DB = db
	fmt.Println("Database connected")
}
