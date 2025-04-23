package initializers

import (
	"log"
	"os"

	"github.com/bryx-ltd/user-service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnection() {
	var err error

	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database")
	}
}

func SyncDatabase() {
	DB.AutoMigrate(
		&models.User{})
}
