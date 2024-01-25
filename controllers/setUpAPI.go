package controllers

import (
	"backend_01/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetUpAPI(r *gin.Engine) {

	r.GET("/ping", PingHandler)
	// signUp
	r.POST("/signUp", SignUpHandlerfunc)
	// login
	r.POST("/login", LoginHandler)

	// create a group that requires authentication
	authorized := r.Group("/auth")
	{
		authorized.Use(middleware.AuthMiddleware())
		authorized.GET("/hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"data": "hello"})
		})

		authorized.POST("/createPost", middleware.RoleMiddleware("ADMIN"), CreatePostHandler)
	}
}

func CreatePostHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "create post",
	})
}
