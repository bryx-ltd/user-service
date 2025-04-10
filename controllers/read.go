package controllers

import (
	"github.com/bryx-ltd/user-service/initializers"
	"github.com/bryx-ltd/user-service/models"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	// Get all users
	var users []models.User

	initializers.DB.Find(&users)

	// Return Response
	c.JSON(200, gin.H{
		"Users": users,
	})
}

func GetUserByID(c *gin.Context) {
	// Get User ID
	var user models.User
	id := c.Param("id")

	initializers.DB.First(&user, id)

	// Return Response
	c.JSON(200, gin.H{
		"Users": user,
	})
}
