package controllers

import (
	"github.com/bryx-ltd/user-service/initializers"
	"github.com/bryx-ltd/user-service/models"
	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context) {
	// Get User ID
	id := c.Param("id")

	// Find User to Delete
	var user []models.User
	initializers.DB.Find(&user, id)

	// Delete User
	initializers.DB.Delete(&user)

	// Return Response
	c.JSON(200, gin.H{
		"User": user,
	})
}
