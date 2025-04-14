package main

import (
	"github.com/bryx-ltd/user-service/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadDotEnv()
	initializers.DBConnection()
}

func main() {
	router := gin.Default()

	router.POST("/users", PostUser)
	router.PUT("/users/:id", UpdateUser)
	router.GET("/users", GetUsers)
	router.GET("/users/:id", GetUserByID)
	router.DELETE("/users/:id", DeleteUser)

	router.Run()
}
