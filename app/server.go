package app

import (
	"majoo-backend-test/app/routes"
	"majoo-backend-test/app/utils"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Create Init App Function
func InitApp() {
	// Setup Database Connection
	db := utils.SetupDatabase()

	// Get Gin Mode From ENV
	mode := os.Getenv("GIN_MODE")

	// Set Gin Mode
	gin.SetMode(mode)

	// Create New App Instance
	app := gin.Default()

	// Setup CORS
	app.Use(cors.Default())

	// Setup Routes
	routes.SetupRoutes(app, db)

	// Run App at 8080
	app.Run(":8080")
}
