package controllers

import (
	"github.com/bryx-ltd/user-service/initializers"
	"github.com/bryx-ltd/user-service/models"
	"github.com/gin-gonic/gin"
)

func PostUser(c *gin.Context) {
	var input struct {
		FirstName    string
		LastName     string
		EmailAddress string
		Password     string
	}
	c.Bind(&input)

	user := models.User{FirstName: input.FirstName, LastName: input.LastName, EmailAddress: input.EmailAddress, Password: input.Password}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"User": user,
	})
}

func GetUsers(c *gin.Context) {
	var users []models.User

	initializers.DB.Find(&users)

	c.JSON(200, gin.H{
		"Users": users,
	})
}

func GetUserByID(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	initializers.DB.First(&user, id)

	c.JSON(200, gin.H{
		"Users": user,
	})
}
