package controllers

import (
	"github.com/bryx-ltd/user-service/initializers"
	"github.com/bryx-ltd/user-service/models"
	"github.com/gin-gonic/gin"
)

func UpdateUser(c *gin.Context) {
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
