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

// Create New Outlet Controller Inteface
type OutletController interface {
	CreateOutlet(*gin.Context)
	UpdateOutlet(*gin.Context)
	DeleteOutlet(*gin.Context)
	GetListOutletByUserId(*gin.Context)
	GetListOutletByUserIdAndMerchantId(*gin.Context)
	GetOutletDetailByUserIdAndOutletId(*gin.Context)
}

// Binding Service to Controller ( Constructor )
type outletController struct {
	service         services.OutletService
	merchantService services.MerchantService
}

// Create Function, to Create New Outlet Controller Instance
func NewOutletController(service services.OutletService, merchantService services.MerchantService) *outletController {
	return &outletController{
		service:         service,
		merchantService: merchantService,
	}
}

///
/// Implementing All Outlet Controller Interface
///

// Create Outlet
func (controller *outletController) CreateOutlet(context *gin.Context) {
	// Get User Detail from JWT
	user, err := utils.GetUserDetailFromToken(context)

	// Check if there is error when get user detail
	if err != nil {
		common.SendError(context, http.StatusUnauthorized, 5, "Invalid token", []string{err.Error()})
		return
	}

	// Create New Outlet Request
	var request models.OutletRequest

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

	// Check if Merchant Exist
	isMerchantExist, err := controller.merchantService.CheckIfMerchantExistByMerchantId(request.MerchantId)

	// Check error
	if !isMerchantExist {
		// If Query Error
		if err != nil {
			common.SendError(context, http.StatusInternalServerError, 1, "Internal server error", []string{err.Error()})
			return
		}

		common.SendError(context, http.StatusBadRequest, 1, "Merchant not found", []string{"merchant: merchant not found"})
		return
	}

	// Access Service to Create Outlet
	errrorCreate := controller.service.CreateOutlet(request, user.Id)

	// Check if there is error when Create Outlet
	if errrorCreate != nil {
		common.SendError(context, http.StatusInternalServerError, 1, "Internal server error", []string{errrorCreate.Error()})
		return
	}

	// Return Success Create Outlet
	common.SendSuccess(context, http.StatusCreated, 0, "Create Outlet Success", nil)
}

// Update Outlet
func (controller *outletController) UpdateOutlet(context *gin.Context) {
	// Event Id
	outletId, _ := strconv.Atoi(context.Param("id"))

	// Access Service to Get Outlet Detail
	outlet, err := controller.service.GetOutletByOutletId(outletId)

	// Check error
	if err != nil {
		common.SendError(context, http.StatusBadRequest, 1, "Outlet not found", []string{"outlet: outlet not found"})
		return
	}

	// Get User Detail from JWT
	user, err := utils.GetUserDetailFromToken(context)

	// Check if there is error when get user detail
	if err != nil {
		common.SendError(context, http.StatusUnauthorized, 5, "Invalid token", []string{err.Error()})
		return
	}

	// Check if User own related Outlet
	if outlet.UserId != user.Id {
		common.SendError(context, http.StatusBadRequest, 1, "You don't have permission to update this outlet", []string{"outlet: You don't have permission to update this outlet"})
		return
	}

	// Create New Outlet Request
	var request models.OutletRequest

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

	// Access Service to Create Outlet
	errorUpdate := controller.service.UpdateOutlet(request, outletId)

	// Check if there is error when create outlet
	if errorUpdate != nil {
		common.SendError(context, http.StatusInternalServerError, 1, "Internal server errorUpdateor", []string{errorUpdate.Error()})
		return
	}

	// Return Success Create Outlet
	common.SendSuccess(context, http.StatusOK, 0, "Update Outlet Success", nil)
}

// Delete Outlet
func (controller *outletController) DeleteOutlet(context *gin.Context) {
	// Outlet Id
	outletId, _ := strconv.Atoi(context.Param("id"))

	// Access Service to Get Outlet Detail
	outlet, err := controller.service.GetOutletByOutletId(outletId)

	// Check error
	if err != nil {
		common.SendError(context, http.StatusBadRequest, 1, "Outlet not found", []string{"outlet: outlet not found"})
		return
	}

	// Get User Detail from JWT
	user, err := utils.GetUserDetailFromToken(context)

	// Check if there is error when get user detail
	if err != nil {
		common.SendError(context, http.StatusUnauthorized, 5, "Invalid token", []string{err.Error()})
		return
	}

	// Check if User own related Outlet
	if outlet.UserId != user.Id {
		common.SendError(context, http.StatusBadRequest, 1, "You don't have permission to update this outlet", []string{"outlet: You don't have permission to update this outlet"})
		return
	}

	// Access Service to Delete Outlet
	errorDelete := controller.service.DeleteOutlet(outletId)

	// Check if there is error when delete outlet
	if errorDelete != nil {
		common.SendError(context, http.StatusInternalServerError, 1, "Internal server error", []string{errorDelete.Error()})
		return
	}

	// Return Success Delete Outlet
	common.SendSuccess(context, http.StatusOK, 0, "Delete Outlet Success", nil)
}

// Get List Outlet By User Id
func (controller *outletController) GetListOutletByUserId(context *gin.Context) {
	// Get User Detail from JWT
	user, err := utils.GetUserDetailFromToken(context)

	// Check if there is error when get user detail
	if err != nil {
		common.SendError(context, http.StatusUnauthorized, 5, "Invalid token", []string{err.Error()})
		return
	}

	// Access Service to Get List Outlet By User Id
	listOutlet, errorGetList := controller.service.GetListOutletByUserId(user.Id)

	// Check if there is error when get list outlet by user id from database
	if errorGetList != nil {
		common.SendError(context, http.StatusInternalServerError, 1, "Internal server error", []string{errorGetList.Error()})
		return
	}

	// Return Success Get List Outlet By User Id
	common.SendSuccess(context, http.StatusOK, 0, "Get List Outlet Success", listOutlet)
}

// Get List Outlet By User Id
func (controller *outletController) GetListOutletByUserIdAndMerchantId(context *gin.Context) {
	// Outlet Id
	merchantId, _ := strconv.Atoi(context.Param("id"))

	// Get User Detail from JWT
	user, err := utils.GetUserDetailFromToken(context)

	// Check if there is error when get user detail
	if err != nil {
		common.SendError(context, http.StatusUnauthorized, 5, "Invalid token", []string{err.Error()})
		return
	}

	// Access Service to Get List Outlet By User Id
	listOutlet, errorGetList := controller.service.GetListOutletByUserIdAndMerchantId(user.Id, merchantId)

	// Check if there is error when get list mechant by user id from database
	if errorGetList != nil {
		common.SendError(context, http.StatusInternalServerError, 1, "Internal server error", []string{errorGetList.Error()})
		return
	}

	// Return Success Get List Outlet By User Id
	common.SendSuccess(context, http.StatusOK, 0, "Get List Outlet Success", listOutlet)
}

// Get Outlet Detail By User Id & OutletId
func (controller *outletController) GetOutletDetailByUserIdAndOutletId(context *gin.Context) {
	// Outlet Id
	outletId, _ := strconv.Atoi(context.Param("id"))

	// Get User Detail from JWT
	user, err := utils.GetUserDetailFromToken(context)

	// Check if there is error when get user detail
	if err != nil {
		common.SendError(context, http.StatusUnauthorized, 5, "Invalid token", []string{err.Error()})
		return
	}

	// Access Service to Get Outlet Detail By User Id & OutletId
	outlet, errorGet := controller.service.GetOutletDetailByUserIdAndOutletId(user.Id, outletId)

	// Check if there is error when get outlet detail by user id & outlet id from database
	if errorGet != nil {
		// If Outlet not found
		if errorGet == sql.ErrNoRows {
			common.SendError(context, http.StatusBadRequest, 1, "Outlet not found", []string{"outlet: outlet id not found"})
			return
		}

		common.SendError(context, http.StatusInternalServerError, 1, "Internal server error", []string{errorGet.Error()})
		return
	}

	// Return Success Get Outlet Detail By User Id & OutletId
	common.SendSuccess(context, http.StatusOK, 0, "Get List Outlet Success", outlet)
}
