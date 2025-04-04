package main

import (
	"github.com/bryx-ltd/user-service/initializers"
	"github.com/bryx-ltd/user-service/models"
)

func init() {
	initializers.LoadDotEnv()
	initializers.DBConnection()
}

// AutoMigrate will create tables, missing foreign keys, constraints, columns and indexes.
func main() {
	initializers.DB.AutoMigrate(
		&models.User{})
}
