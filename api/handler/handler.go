package handler

import (
	"fmt"
	"iman_task/helper/jwt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func GenerateToken(c *gin.Context) {
	token, err := jwt.GenerateToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func JwtMiddleware(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code" : http.StatusUnauthorized,
			"error": "Token required"})
		c.Abort()
		return
	}

	tokens := strings.Split(tokenString, " ")
	if len(tokens) != 2 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code" : http.StatusUnauthorized,
			"error": "Invalid token"})
		c.Abort()
		return
	}

	tokenString = tokens[1]
	claims, err := jwt.VerifyToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code" : http.StatusUnauthorized,
			"error": "Invalid token"})
		c.Abort()
		return
	}

	c.Set("claims", claims)
	c.Next()
}

func GetDaysLeft(c *gin.Context) {
	currentTime := time.Now()
	targetTime := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)
	remainingDuration := targetTime.Sub(currentTime)
	daysLeft := int(remainingDuration.Hours() / 24)

	c.JSON(http.StatusOK, gin.H{
		"daysLeft": daysLeft,
	})

	fmt.Println("Days left till 2025-01-01: ", daysLeft)
}


