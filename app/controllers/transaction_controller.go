package controllers

import (
	"majoo-backend-test/app/common"
	"majoo-backend-test/app/services"
	"majoo-backend-test/app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create New Transaction Controller Inteface
type TransactionController interface {
	GetListTransactionReportByUserId(*gin.Context)
	GetListTransactionReportDetailByUserId(*gin.Context)
}

// Binding Service to Controller ( Constructor )
type transactionController struct {
	service         services.TransactionService
	merchantService services.MerchantService
	outletService   services.OutletService
}

// Create Function, to Create New Transaction Controller Instance
func NewTransactionController(service services.TransactionService, merchantService services.MerchantService, outletService services.OutletService) *transactionController {
	return &transactionController{
		service:         service,
		merchantService: merchantService,
		outletService:   outletService,
	}
}

// Get List Transaction Report By User Id
func (controller *transactionController) GetListTransactionReportByUserId(context *gin.Context) {
	// Get User Detail from JWT
	user, err := utils.GetUserDetailFromToken(context)

	// Check if there is error when get user detail
	if err != nil {
		common.SendError(context, http.StatusUnauthorized, 5, "Invalid token", []string{err.Error()})
		return
	}

	// Access Service to Get List Transaction By User Id
	listTransaction, errorGetList := controller.service.GetListTransactionReportByUserId(user.Id)

	// Check if there is error when get list outlet by user id from database
	if errorGetList != nil {
		common.SendError(context, http.StatusInternalServerError, 1, "Internal server error", []string{errorGetList.Error()})
		return
	}

	// Return Success Get List Transaction By User Id
	common.SendSuccess(context, http.StatusOK, 0, "Get List Transaction Success", listTransaction)
}

// Get List Transaction Report Detail By User Id
func (controller *transactionController) GetListTransactionReportDetailByUserId(context *gin.Context) {
	// Get User Detail from JWT
	user, err := utils.GetUserDetailFromToken(context)

	// Check if there is error when get user detail
	if err != nil {
		common.SendError(context, http.StatusUnauthorized, 5, "Invalid token", []string{err.Error()})
		return
	}

	// Access Service to Get List Transaction By User Id
	listTransaction, errorGetList := controller.service.GetListTransactionReportDetailByUserId(user.Id)

	// Check if there is error when get list outlet by user id from database
	if errorGetList != nil {
		common.SendError(context, http.StatusInternalServerError, 1, "Internal server error", []string{errorGetList.Error()})
		return
	}

	// Return Success Get List Transaction By User Id
	common.SendSuccess(context, http.StatusOK, 0, "Get List Transaction Success", listTransaction)
}
