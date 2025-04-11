package controllers

import (
	"github.com/bryx-ltd/user-service/initializers"
	"github.com/bryx-ltd/user-service/models"
	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context) {
	// Get User ID
	id := c.Param("id")

	// Delete User
	initializers.DB.Delete(&models.User{}, id)

	// Return Response
	c.Status(200)
}
