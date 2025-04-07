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

	router.POST("/user", controllers.PostUser)

	router.Run()
}
