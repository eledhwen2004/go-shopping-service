package postgre

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() error {
	dsn := "host=localhost user=myuser password=mypassword database=go-db port=5434 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error while connecting database: ", err)
		os.Exit(1)
	}

	DB = db

	return nil
}
