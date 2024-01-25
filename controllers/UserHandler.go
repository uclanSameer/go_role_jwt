package controllers

import (
	"backend_01/models"
	"backend_01/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// LoginHandler authenticates the user and returns a JWT token
func LoginHandler(c *gin.Context) {
	var user models.User
	// Bind the request body to the user struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	// Check the user's credentials
	token, err := services.Login(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func SignUpHandlerfunc(ctx *gin.Context) {
	var user models.User
	ctx.BindJSON(&user)
	services.CreateUser(&user)
	ctx.JSON(200, gin.H{
		"message": "success",
	})
}
