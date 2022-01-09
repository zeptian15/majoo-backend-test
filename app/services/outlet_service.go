package services

import (
	"majoo-backend-test/app/models"
	"majoo-backend-test/app/repositories"
)

// Create New Outlet Service Interface
type OutletService interface {
	CreateOutlet(outlet models.OutletRequest, userId int) error
	UpdateOutlet(outlet models.OutletRequest, outletId int) error
	DeleteOutlet(outletId int) error
	GetOutletByOutletId(outletId int) (models.OutletResponse, error)
	GetListOutletByUserId(userId int) ([]models.OutletResponse, error)
	GetListOutletByUserIdAndMerchantId(userId int, merchantId int) ([]models.OutletResponse, error)
	GetOutletDetailByUserIdAndOutletId(userId int, outletId int) (models.OutletResponse, error)
	CheckIfOutletExistByOutletId(outletId int) (bool, error)
}

// Binding Repository to Service ( Constructor )
type outletService struct {
	repository repositories.OutletRepository
}

// Create New Function, to Create New Outlet Service Instance
func NewOutletService(repository repositories.OutletRepository) *outletService {
	return &outletService{repository}
}

///
/// Implementing Outlet Service Interface
///

// Create Outlet
func (service *outletService) CreateOutlet(outletRequest models.OutletRequest, userId int) error {
	// Convert Outlet Request to Outlet Model
	var outlet models.Outlet
	outlet.UserId = userId
	outlet.MerchantId = outletRequest.MerchantId
	outlet.OutletName = outletRequest.OutletName

	// Access Repository to Create Outlet
	err := service.repository.CreateOutlet(outlet)

	// Return Error
	return err
}

// Update Outlet
func (service *outletService) UpdateOutlet(outletRequest models.OutletRequest, outletId int) error {
	// Convert Outlet Request to Outlet Model
	var outlet models.Outlet
	outlet.MerchantId = outletRequest.MerchantId
	outlet.OutletName = outletRequest.OutletName

	// Access Repository to Update Outlet
	err := service.repository.UpdateOutlet(outlet, outletId)

	// Return Error
	return err
}

// Delete Outlet
func (service *outletService) DeleteOutlet(outletId int) error {
	// Access Repository to Delete Outlet
	err := service.repository.DeleteOutlet(outletId)

	// Return Error
	return err
}

// Get Outlet Detail by Outlet Id
func (service *outletService) GetOutletByOutletId(outletId int) (models.OutletResponse, error) {
	// Access Repository to Get Outlet Detail by Outlet Id
	outlet, err := service.repository.GetOutletByOutletId(outletId)

	// Return Error
	return outlet, err
}

// Get List Outlet by User Id
func (service *outletService) GetListOutletByUserId(userId int) ([]models.OutletResponse, error) {
	// Access Repository to Get List Outlet by User Id
	listOutlet, err := service.repository.GetListOutletByUserId(userId)

	// Return Error
	return listOutlet, err
}

// Get List Outlet by User Id & Merchant Id
func (service *outletService) GetListOutletByUserIdAndMerchantId(userId int, merchantId int) ([]models.OutletResponse, error) {
	// Access Repository to Get List Outlet by User Id
	listOutlet, err := service.repository.GetListOutletByUserIdAndMerchantId(userId, merchantId)

	// Return Error
	return listOutlet, err
}

// Get Outlet Detail by User Id & Outlet Id
func (service *outletService) GetOutletDetailByUserIdAndOutletId(userId int, outletId int) (models.OutletResponse, error) {
	// Access Repository to Get Outlet Detail by User Id & Outlet Id
	outlet, err := service.repository.GetOutletDetailByUserIdAndOutletId(userId, outletId)

	// Return Error
	return outlet, err
}

// Check if Outlet Exists by Outlet Id
func (service *outletService) CheckIfOutletExistByOutletId(outletId int) (bool, error) {
	// Access Repository to Check if Outlet Exists by Outlet Id
	isExist, err := service.repository.CheckIfOutletExistByOutletId(outletId)

	// Return Error
	return isExist, err
}
