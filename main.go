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

	router.POST("/users", handleCreateUser)
	router.PUT("/users/:id", handleUpdateUser)
	router.GET("/users", handleGetUsers)
	router.GET("/users/:id", handleGetUserByID)
	router.DELETE("/users/:id", handleDeleteUser)

	router.Run()
}
