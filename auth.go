package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/bryx-ltd/user-service/initializers"
	"github.com/bryx-ltd/user-service/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func authorizeUser(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization cookie not found"})
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		return
	}

	if exp, ok := claims["exp"].(float64); !ok || float64(time.Now().Unix()) > exp {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token has expired"})
		return
	}

	var user models.User
	if err := initializers.DB.First(&user, claims["sub"]).Error; err != nil || user.ID == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	c.Set("user", user)
	c.Next()
}
