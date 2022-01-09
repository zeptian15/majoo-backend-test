package services

import (
	"majoo-backend-test/app/models"
	"majoo-backend-test/app/repositories"
)

// Create New Merchant Service Interface
type MerchantService interface {
	CreateMerchant(merchant models.MerchantRequest, userId int) error
	UpdateMerchant(merchant models.MerchantRequest, merchantId int) error
	DeleteMerchant(merchantId int) error
	GetMerchantByMerchantId(merchantId int) (models.MerchantResponse, error)
	GetListMerchantByUserId(userId int) ([]models.MerchantResponse, error)
	GetMerchantDetailByUserIdAndMerchantId(userId int, merchantId int) (models.MerchantResponse, error)
	CheckIfMerchantExistByMerchantId(merchantId int) (bool, error)
}

// Binding Repository to Service ( Constructor )
type merchantService struct {
	repository repositories.MerchantRepository
}

// Create New Function, to Create New Merchant Service Instance
func NewMerchantService(repository repositories.MerchantRepository) *merchantService {
	return &merchantService{repository}
}

///
/// Implementing Merchant Service Interface
///

// Create Merchant
func (service *merchantService) CreateMerchant(merchantRequest models.MerchantRequest, userId int) error {
	// Convert Merchant Request to Merchant Model
	var merchant models.Merchant
	merchant.UserId = userId
	merchant.MerchantName = merchantRequest.MerchantName

	// Access Repository to Create Merchant
	err := service.repository.CreateMerchant(merchant)

	// Return Error
	return err
}

// Update Merchant
func (service *merchantService) UpdateMerchant(merchantRequest models.MerchantRequest, merchantId int) error {
	// Convert Merchant Request to Merchant Model
	var merchant models.Merchant
	merchant.MerchantName = merchantRequest.MerchantName

	// Access Repository to Update Merchant
	err := service.repository.UpdateMerchant(merchant, merchantId)

	// Return Error
	return err
}

// Delete Merchant
func (service *merchantService) DeleteMerchant(merchantId int) error {
	// Access Repository to Delete Merchant
	err := service.repository.DeleteMerchant(merchantId)

	// Return Error
	return err
}

// Get Merchant Detail by Merchant Id
func (service *merchantService) GetMerchantByMerchantId(merchantId int) (models.MerchantResponse, error) {
	// Access Repository to Get Merchant Detail by Merchant Id
	merchant, err := service.repository.GetMerchantByMerchantId(merchantId)

	// Return Error
	return merchant, err
}

// Get List Merchant by User Id
func (service *merchantService) GetListMerchantByUserId(userId int) ([]models.MerchantResponse, error) {
	// Access Repository to Get List Merchant by User Id
	listMerchant, err := service.repository.GetListMerchantByUserId(userId)

	// Return Error
	return listMerchant, err
}

// Get Merchant Detail by User Id & Merchant Id
func (service *merchantService) GetMerchantDetailByUserIdAndMerchantId(userId int, merchantId int) (models.MerchantResponse, error) {
	// Access Repository to Get Merchant Detail by User Id & Merchant Id
	merchant, err := service.repository.GetMerchantDetailByUserIdAndMerchantId(userId, merchantId)

	// Return Error
	return merchant, err
}

// Check if Merchant Exists by Merchant Id
func (service *merchantService) CheckIfMerchantExistByMerchantId(merchantId int) (bool, error) {
	// Access Repository to Check if Merchant Exists by Merchant Id
	isExist, err := service.repository.CheckIfMerchantExistByMerchantId(merchantId)

	// Return Error
	return isExist, err
}
