package routes

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create Setup Routes Function
func SetupRoutes(app *gin.Engine, db *sql.DB) {
	// Check API Health
	app.GET("/", func(context *gin.Context) {
		// Return Message
		context.JSON(http.StatusOK, gin.H{
			"message": "Server Alive!",
		})
	})
}
