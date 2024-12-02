package main

import (
	"api-gateway/src/routes"
	"api-gateway/src/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load service configurations
	config := utils.LoadConfig()

	// Initialize Gin router
	router := gin.Default()

	// Register routes
	routes.RegisterRoutes(router, config.Services)

	// Start API Gateway
	router.Run(":8080") // Gateway listens on port 8080
}
