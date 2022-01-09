package controllers

import (
	"database/sql"
	"majoo-backend-test/app/common"
	"majoo-backend-test/app/models"
	"majoo-backend-test/app/services"
	"majoo-backend-test/app/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Create New Merchant Controller Inteface
type MerchantController interface {
	CreateMerchant(*gin.Context)
	UpdateMerchant(*gin.Context)
	DeleteMerchant(*gin.Context)
	GetListMerchantByUserId(*gin.Context)
	GetMerchantDetailByUserIdAndMerchantId(*gin.Context)
}

// Binding Service to Controller ( Constructor )
type merchantController struct {
	service services.MerchantService
}

// Create Function, to Create New Merchant Controller Instance
func NewMerchantController(service services.MerchantService) *merchantController {
	return &merchantController{service}
}

///
/// Implementing All Merchant Controller Interface
///

// Create Merchant
func (controller *merchantController) CreateMerchant(context *gin.Context) {
	// Get User Detail from JWT
	user, err := utils.GetUserDetailFromToken(context)

	// Check if there is error when get user detail
	if err != nil {
		common.SendError(context, http.StatusUnauthorized, 5, "Invalid token", []string{err.Error()})
		return
	}

	// Create New Merchant Request
	var request models.MerchantRequest

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

	// Access Service to Create Merchant
	errrorCreate := controller.service.CreateMerchant(request, user.Id)

	// Check if there is error when Create Merchant
	if errrorCreate != nil {
		common.SendError(context, http.StatusInternalServerError, 1, "Internal server error", []string{errrorCreate.Error()})
		return
	}

	// Return Success Create Merchant
	common.SendSuccess(context, http.StatusCreated, 0, "Create Merchant Success", nil)
}

// Update Merchant
func (controller *merchantController) UpdateMerchant(context *gin.Context) {
	// Event Id
	merchantId, _ := strconv.Atoi(context.Param("id"))

	// Access Service to Get Merchant Detail
	merchant, err := controller.service.GetMerchantByMerchantId(merchantId)

	// Check error
	if err != nil {
		common.SendError(context, http.StatusBadRequest, 1, "Merchant not found", []string{"merchant: merchant not found"})
		return
	}

	// Get User Detail from JWT
	user, err := utils.GetUserDetailFromToken(context)

	// Check if there is error when get user detail
	if err != nil {
		common.SendError(context, http.StatusUnauthorized, 5, "Invalid token", []string{err.Error()})
		return
	}

	// Check if User own related Merchant
	if merchant.UserId != user.Id {
		common.SendError(context, http.StatusBadRequest, 1, "You don't have permission to update this merchant", []string{"merchant: You don't have permission to update this merchant"})
		return
	}

	// Create New Merchant Request
	var request models.MerchantRequest

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

	// Access Service to Create Merchant
	errorUpdate := controller.service.UpdateMerchant(request, merchantId)

	// Check if there is error when create merchant
	if errorUpdate != nil {
		common.SendError(context, http.StatusInternalServerError, 1, "Internal server errorUpdateor", []string{errorUpdate.Error()})
		return
	}

	// Return Success Create Merchant
	common.SendSuccess(context, http.StatusOK, 0, "Update Merchant Success", nil)
}

// Delete Merchant
func (controller *merchantController) DeleteMerchant(context *gin.Context) {
	// Merchant Id
	merchantId, _ := strconv.Atoi(context.Param("id"))

	// Access Service to Get Merchant Detail
	merchant, err := controller.service.GetMerchantByMerchantId(merchantId)

	// Check error
	if err != nil {
		common.SendError(context, http.StatusBadRequest, 1, "Merchant not found", []string{"merchant: merchant not found"})
		return
	}

	// Get User Detail from JWT
	user, err := utils.GetUserDetailFromToken(context)

	// Check if there is error when get user detail
	if err != nil {
		common.SendError(context, http.StatusUnauthorized, 5, "Invalid token", []string{err.Error()})
		return
	}

	// Check if User own related Merchant
	if merchant.UserId != user.Id {
		common.SendError(context, http.StatusBadRequest, 1, "You don't have permission to update this merchant", []string{"merchant: You don't have permission to update this merchant"})
		return
	}

	// Access Service to Delete Merchant
	errorDelete := controller.service.DeleteMerchant(merchantId)

	// Check if there is error when delete merchant
	if errorDelete != nil {
		common.SendError(context, http.StatusInternalServerError, 1, "Internal server error", []string{errorDelete.Error()})
		return
	}

	// Return Success Delete Merchant
	common.SendSuccess(context, http.StatusOK, 0, "Delete Merchant Success", nil)
}

// Get List Merchant By User Id
func (controller *merchantController) GetListMerchantByUserId(context *gin.Context) {
	// Get User Detail from JWT
	user, err := utils.GetUserDetailFromToken(context)

	// Check if there is error when get user detail
	if err != nil {
		common.SendError(context, http.StatusUnauthorized, 5, "Invalid token", []string{err.Error()})
		return
	}

	// Access Service to Get List Merchant By User Id
	listMerchant, errorGetList := controller.service.GetListMerchantByUserId(user.Id)

	// Check if there is error when get list mechant by user id from database
	if errorGetList != nil {
		common.SendError(context, http.StatusInternalServerError, 1, "Internal server error", []string{errorGetList.Error()})
		return
	}

	// Return Success Get List Merchant By User Id
	common.SendSuccess(context, http.StatusOK, 0, "Get List Merchant Success", listMerchant)
}

// Get Merchant Detail By User Id & MerchantId
func (controller *merchantController) GetMerchantDetailByUserIdAndMerchantId(context *gin.Context) {
	// Merchant Id
	merchantId, _ := strconv.Atoi(context.Param("id"))

	// Get User Detail from JWT
	user, err := utils.GetUserDetailFromToken(context)

	// Check if there is error when get user detail
	if err != nil {
		common.SendError(context, http.StatusUnauthorized, 5, "Invalid token", []string{err.Error()})
		return
	}

	// Access Service to Get Merchant Detail By User Id & MerchantId
	merchant, errorGet := controller.service.GetMerchantDetailByUserIdAndMerchantId(user.Id, merchantId)

	// Check if there is error when get mechant detail by user id & merchant id from database
	if errorGet != nil {
		// If Merchant not found
		if errorGet == sql.ErrNoRows {
			common.SendError(context, http.StatusBadRequest, 1, "Merchant not found", []string{"merchant: merchant id not found"})
			return
		}

		common.SendError(context, http.StatusInternalServerError, 1, "Internal server error", []string{errorGet.Error()})
		return
	}

	// Return Success Get Merchant Detail By User Id & MerchantId
	common.SendSuccess(context, http.StatusOK, 0, "Get List Merchant Success", merchant)
}
