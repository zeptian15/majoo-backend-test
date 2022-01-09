package services

import (
	"majoo-backend-test/app/models"
	"majoo-backend-test/app/repositories"
)

// Create New Transaction Service Interface
type TransactionService interface {
	GetListTransactionReportByUserId(userId int) ([]models.TransactionResponse, error)
	GetListTransactionReportDetailByUserId(userId int) ([]models.TransactionDetailResponse, error)
}

// Binding Repository to Service ( Constructor )
type transactionService struct {
	repository repositories.TransactionRepository
}

// Create New Function, to Create New Transaction Service Instance
func NewTransactionService(repository repositories.TransactionRepository) *transactionService {
	return &transactionService{repository}
}

///
/// Implementing Transaction Service Interface
///

// Get List Transaction Report By User Id
func (service *transactionService) GetListTransactionReportByUserId(userId int) ([]models.TransactionResponse, error) {
	// Access Repository to Get Transaction Detail by Transaction Id
	transaction, err := service.repository.GetListTransactionReportByUserId(userId)

	// Return Error
	return transaction, err
}

// Get List Transaction Report Detail By User Id
func (service *transactionService) GetListTransactionReportDetailByUserId(userId int) ([]models.TransactionDetailResponse, error) {
	// Access Repository to Get Transaction Detail by Transaction Id
	transaction, err := service.repository.GetListTransactionReportDetailByUserId(userId)

	// Return Error
	return transaction, err
}
