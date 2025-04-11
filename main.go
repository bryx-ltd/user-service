package main

import (
	"github.com/bryx-ltd/user-service/controllers"
	"github.com/bryx-ltd/user-service/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadDotEnv()
	initializers.DBConnection()
}

func main() {
	router := gin.Default()

	router.POST("/users", controllers.PostUser)
	router.PUT("/users/:id", controllers.UpdateUser)
	router.GET("/users", controllers.GetUsers)
	router.GET("/users/:id", controllers.GetUserByID)
	router.DELETE("/users/:id", controllers.DeleteUser)

	router.Run()
}
