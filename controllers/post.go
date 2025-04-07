package controllers

import (
	"github.com/bryx-ltd/user-service/initializers"
	"github.com/bryx-ltd/user-service/models"
	"github.com/gin-gonic/gin"
)

func PostUser(c *gin.Context) {

	// Create a hardcoded user
	user := models.User{FirstName: "Neil", LastName: "Davies", EmailAddress: "test@test.com", Password: "verySecurePassword"}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"User": user,
	})
}
