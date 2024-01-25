package main

import (
	"backend_01/config"
	"backend_01/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	client := config.InitDataSource()
	defer client.Close()

	r := gin.Default()

	// Set up API
	controllers.SetUpAPI(r)
	r.Run()
}
