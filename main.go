package main

import (
	"github.com/bryx-ltd/user-service/database"
	"github.com/bryx-ltd/user-service/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadDotEnv()
	database.DBConnection()
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "success",
		})
	})

	router.Run()
}
