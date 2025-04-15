package main

import (
	"net/http"

	"github.com/bryx-ltd/user-service/initializers"
	"github.com/bryx-ltd/user-service/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func handleCreateUser(c *gin.Context) {
	// Get Request Data
	var body struct {
		FirstName    string
		LastName     string
		EmailAddress string
		Password     string
	}

	// Bind Request Data to Struct
	// This will automatically parse the JSON body and bind it to the struct
	// If the JSON is invalid, it will return a 400 Bad Request response
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Invalid request data",
		})
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to hash password",
		})
		return
	}

	// Create User
	user := models.User{FirstName: body.FirstName, LastName: body.LastName, EmailAddress: body.EmailAddress, Password: string(hashedPassword)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to create user",
		})
		return
	}

	// Return Response
	c.JSON(http.StatusOK, gin.H{
		"User": user,
	})
}

func handleGetUsers(c *gin.Context) {
	// Get all users
	var users []models.User

	initializers.DB.Find(&users)

	// Return Response
	c.JSON(http.StatusOK, gin.H{
		"Users": users,
	})
}

func handleGetUserByID(c *gin.Context) {
	// Get User ID
	var user models.User
	id := c.Param("id")

	initializers.DB.First(&user, id)

	// Return Response
	c.JSON(http.StatusOK, gin.H{
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

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Invalid request data",
		})
		return
	}

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
	c.Status(http.StatusOK)
}
