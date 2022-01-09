package routes

import (
	"database/sql"
	"majoo-backend-test/app/controllers"
	"majoo-backend-test/app/middlewares"
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
	userRepository := repositories.NewUserRepository(db)
	authService := services.NewAuthService(userRepository)
	authController := controllers.NewAuthController(authService)

	/// Merchant ///
	merchantRepository := repositories.NewMerchantRepository(db)
	merchantService := services.NewMerchantService(merchantRepository)
	merchantController := controllers.NewMerchantController(merchantService)

	/// Outlet ///
	outletRepository := repositories.NewOutletRepository(db)
	outletService := services.NewOutletService(outletRepository)
	outletController := controllers.NewOutletController(outletService, merchantService)

	// Setup Routes Group
	v1 := app.Group("/api/v1")

	/// Middlewares ///
	authMiddleware := middlewares.TokenAuthMiddleware(userRepository)

	// Setup Routes

	/// Auth Routes ///
	v1.POST("/auth/register", authController.RegisterUser)
	v1.POST("/auth/login", authController.LoginUser)
	v1.GET("/auth/profile", authController.GetUserProfile)

	/// Merchants Routes ///
	v1.POST("/merchants", authMiddleware, merchantController.CreateMerchant)
	v1.PATCH("/merchants/:id", authMiddleware, merchantController.UpdateMerchant)
	v1.DELETE("/merchants/:id", authMiddleware, merchantController.DeleteMerchant)
	v1.GET("/merchants", authMiddleware, merchantController.GetListMerchantByUserId)
	v1.GET("/merchants/:id/detail", authMiddleware, merchantController.GetMerchantDetailByUserIdAndMerchantId)

	/// Merchants Routes ///
	v1.POST("/outlets", authMiddleware, outletController.CreateOutlet)
	v1.PATCH("/outlets/:id", authMiddleware, outletController.UpdateOutlet)
	v1.DELETE("/outlets/:id", authMiddleware, outletController.DeleteOutlet)
	v1.GET("/outlets", authMiddleware, outletController.GetListOutletByUserId)
	v1.GET("/outlets/merchants/:id", authMiddleware, outletController.GetListOutletByUserIdAndMerchantId)
	v1.GET("/outlets/:id/detail", authMiddleware, outletController.GetOutletDetailByUserIdAndOutletId)
}
