package routes

import (
	"database/sql"
	"majoo-backend-test/app/controllers"
	"majoo-backend-test/app/repositories"
	"majoo-backend-test/app/services"
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

	// Binding Required Controller, Service, & Repository

	/// Auth - User ///
	authRepository := repositories.NewUserRepository(db)
	authService := services.NewAuthService(authRepository)
	authController := controllers.NewAuthController(authService)

	// Setup Routes Group
	v1 := app.Group("/api/v1")

	// Setup Routes

	/// Auth Routes ///
	v1.POST("/auth/register", authController.RegisterUser)
	v1.POST("/auth/login", authController.LoginUser)
	v1.GET("/auth/profile", authController.GetUserProfile)
}
