package main

import (
	"net/http"

	"github.com/bryx-ltd/user-service/initializers"
	"github.com/bryx-ltd/user-service/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func handleCreateUser(c *gin.Context) {
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to hash password",
		})
		return
	}

	user := models.User{FirstName: body.FirstName, LastName: body.LastName, EmailAddress: body.EmailAddress, Password: string(hashedPassword)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"User": user,
	})
}

func handleGetUsers(c *gin.Context) {
	var users []models.User

	initializers.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{
		"Users": users,
	})
}

func handleGetUserByID(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	initializers.DB.First(&user, id)

	c.JSON(http.StatusOK, gin.H{
		"Users": user,
	})
}

func handleUpdateUser(c *gin.Context) {
	var body struct {
		FirstName    string
		LastName     string
		EmailAddress string
		Password     string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request data",
		})
		return
	}

	id := c.Param("id")

	var user []models.User
	initializers.DB.Find(&user, id)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	initializers.DB.Model(&user).Updates(models.User{FirstName: body.FirstName, LastName: body.LastName, EmailAddress: body.EmailAddress, Password: string(hashedPassword)})

	c.JSON(http.StatusOK, gin.H{
		"users": user,
	})
}

func handleDeleteUser(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.User{}, id)

	c.Status(http.StatusOK)
}
