package main

import (
	"net/http"
	"os"
	"time"

	"github.com/bryx-ltd/user-service/initializers"
	"github.com/bryx-ltd/user-service/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
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

	var user models.User
	initializers.DB.First(&user, "email_address = ?", body.EmailAddress)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "User not found",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Invalid password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to generate token",
		})
		return
	}

	age := 3600 * 24
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, age, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{})
}
