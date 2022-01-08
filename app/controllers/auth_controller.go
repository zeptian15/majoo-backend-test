package controllers

import (
	"majoo-backend-test/app/common"
	"majoo-backend-test/app/models"
	"majoo-backend-test/app/services"
	"majoo-backend-test/app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create New Auth Controller Inteface
type AuthController interface {
	RegisterUser(*gin.Context)
	LoginUser(*gin.Context)
	GetUserProfile(*gin.Context)
}

// Binding Service to Controller ( Constructor )
type authController struct {
	service services.AuthService
}

// Create Function, to Create New Auth Controller Instance
func NewAuthController(service services.AuthService) *authController {
	return &authController{service}
}

///
/// Implementing All Auth Controller Interface
///

// Register User
func (controller *authController) RegisterUser(context *gin.Context) {
	// Create New Register Request
	var request models.RegisterRequest

	// Bind body to Model
	errorBinding := context.ShouldBindJSON(&request)

	// Check if there is error when binding from JSON
	if errorBinding != nil {
		// If EOF error
		if errorBinding.Error() == "EOF" {
			common.SendError(context, http.StatusBadRequest, 1, "Body is empty", []string{"Body required"})
			return
		}

		// If Binding Error
		common.SendError(context, http.StatusBadRequest, 1, "Not valid request", utils.SplitError(errorBinding))
		return
	}

	// Access Service to Register User
	err := controller.service.RegisterUser(request)

	// Check if there is error when Register User
	if err != nil {
		common.SendError(context, http.StatusInternalServerError, 1, "Internal server error", []string{err.Error()})
		return
	}

	// Return Success Register
	common.SendSuccess(context, http.StatusCreated, 0, "Register success", nil)
}

// Login User
func (controller *authController) LoginUser(context *gin.Context) {
	// Create New Login Request
	var request models.LoginRequest

	// Bind body to Model
	errorBinding := context.ShouldBindJSON(&request)

	// Check if there is error when binding from JSON
	if errorBinding != nil {
		// If EOF error
		if errorBinding.Error() == "EOF" {
			common.SendError(context, http.StatusBadRequest, 1, "Body is empty", []string{"Body required"})
			return
		}

		// If Binding Error
		common.SendError(context, http.StatusBadRequest, 1, "Not valid request", utils.SplitError(errorBinding))
		return
	}

	// Access Service to Register User
	token, err := controller.service.LoginUser(request)

	// Check if there is error when login
	if err != nil {
		common.SendError(context, http.StatusUnauthorized, 5, "Unathorized", []string{err.Error()})
		return
	}

	// Return Success Login
	common.SendSuccess(context, http.StatusOK, 0, "Login success", gin.H{
		"token": token,
	})
}

// Get User Profile by Authorization Token
func (controller *authController) GetUserProfile(context *gin.Context) {
	// Get User Detail from JWT
	user, err := utils.GetUserDetailFromToken(context)

	// Check if there is error when get user detail
	if err != nil {
		common.SendError(context, http.StatusUnauthorized, 5, "Invalid token", []string{err.Error()})
		return
	}

	// Access Service to Get User Detail
	userDetail, errorDetail := controller.service.GetUserProfile(user.UserName)

	// Check if there is error when get user detail from database
	if errorDetail != nil {
		common.SendError(context, http.StatusInternalServerError, 1, "Internal server error", []string{errorDetail.Error()})
		return
	}

	// Return Success Get Profile User
	common.SendSuccess(context, http.StatusOK, 0, "Get profile success", userDetail)
}
