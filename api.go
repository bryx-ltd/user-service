package main

import (
	"github.com/bryx-ltd/user-service/initializers"
	"github.com/bryx-ltd/user-service/models"
	"github.com/gin-gonic/gin"
)

func handleCreateUser(c *gin.Context) {
	// Get Request Data
	var body struct {
		FirstName    string
		LastName     string
		EmailAddress string
		Password     string
	}
	c.Bind(&body)

	// Create User
	user := models.User{FirstName: body.FirstName, LastName: body.LastName, EmailAddress: body.EmailAddress, Password: body.Password}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return Response
	c.JSON(200, gin.H{
		"User": user,
	})
}

func handleGetUsers(c *gin.Context) {
	// Get all users
	var users []models.User

	initializers.DB.Find(&users)

	// Return Response
	c.JSON(200, gin.H{
		"Users": users,
	})
}

func handleGetUserByID(c *gin.Context) {
	// Get User ID
	var user models.User
	id := c.Param("id")

	initializers.DB.First(&user, id)

	// Return Response
	c.JSON(200, gin.H{
		"Users": user,
	})
}

func handleUpdateUser(c *gin.Context) {
	// Get Request Data
	var body struct {
		FirstName    string
		LastName     string
		EmailAddress string
		Password     string
	}
	c.Bind(&body)

	// Get User ID
	id := c.Param("id")

	//Find User to Update
	var user []models.User
	initializers.DB.Find(&user, id)

	// Update User
	initializers.DB.Model(&user).Updates(models.User{FirstName: body.FirstName, LastName: body.LastName, EmailAddress: body.EmailAddress, Password: body.Password})

	// Return Response
	c.JSON(200, gin.H{
		"Users": user,
	})
}

func handleDeleteUser(c *gin.Context) {
	// Get User ID
	id := c.Param("id")

	// Delete User
	initializers.DB.Delete(&models.User{}, id)

	// Return Response
	c.Status(200)
}
