package controllers

import (
	"github.com/bryx-ltd/user-service/initializers"
	"github.com/bryx-ltd/user-service/models"
	"github.com/gin-gonic/gin"
)

func PostUser(c *gin.Context) {
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
