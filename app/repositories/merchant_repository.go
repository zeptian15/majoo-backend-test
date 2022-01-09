package repositories

import (
	"database/sql"
	"fmt"
	"majoo-backend-test/app/models"
)

// Create New Merchant Repository Interface
type MerchantRepository interface {
	CreateMerchant(models.Merchant) error
	UpdateMerchant(merchant models.Merchant, merchantId int) error
	DeleteMerchant(merchantId int) error
	GetMerchantByMerchantId(merchantId int) (models.MerchantResponse, error)
	GetListMerchantByUserId(userId int) ([]models.MerchantResponse, error)
	GetMerchantDetailByUserIdAndMerchantId(userId int, merchantId int) (models.MerchantResponse, error)
}

// Binding Database to Repository ( Constructor )
type merchantRepository struct {
	db *sql.DB
}

// Create Function, to Create New Merchant Repository Instance
func NewMerchantRepository(db *sql.DB) *merchantRepository {
	return &merchantRepository{db}
}

///
/// Implementing Merchant Repository Interface
///

// Create Merchant
func (repository *merchantRepository) CreateMerchant(merchant models.Merchant) error {
	// Prepare Query
	query := `INSERT INTO merchants ( user_id, merchant_name ) VALUES ( ?, ? )`

	// Execute Query
	_, err := repository.db.Exec(query, merchant.UserId, merchant.MerchantName)

	// Return err
	return err
}

// Update Merchant
func (repository *merchantRepository) UpdateMerchant(merchant models.Merchant, merchantId int) error {
	// Prepare Query
	query := `UPDATE merchants SET merchant_name=? WHERE id=?`

	fmt.Println(merchant.MerchantName, merchant)

	// Execute Query
	_, err := repository.db.Exec(query, merchant.MerchantName, merchantId)

	// Return err
	return err
}

// Delete Merchant ( Soft Detelte )
func (repository *merchantRepository) DeleteMerchant(merchantId int) error {
	// Prepare Query
	query := `UPDATE merchants SET deleted_at=NOW() WHERE id=?`

	// Execute Query
	_, err := repository.db.Exec(query, merchantId)

	// Return err
	return err
}

// Get Merchant by Merchant Id
func (repository *merchantRepository) GetMerchantByMerchantId(merchantId int) (models.MerchantResponse, error) {
	// Prepare Query
	query := `SELECT * FROM merchants WHERE id=? AND deleted_at IS NULL`

	// Prepare New Struct Instance
	var merchant models.MerchantResponse

	// Execute Query
	row := repository.db.QueryRow(query, merchantId)

	// Scan Row to User Response Model
	if err := row.Scan(
		&merchant.Id,
		&merchant.UserId,
		&merchant.MerchantName,
		&merchant.CreatedAt,
		&merchant.UpdatedAt,
		&merchant.DeletedAt,
	); err != nil {
		return models.MerchantResponse{}, err
	}

	// Return User Response Model
	return merchant, nil
}

// Get List Merchant by User Id
func (repository *merchantRepository) GetListMerchantByUserId(userId int) ([]models.MerchantResponse, error) {
	// Prepare Query
	query := `SELECT * FROM merchants WHERE user_id=? AND deleted_at IS NULL`

	// Create List Merchant Response Variable
	var listMerchant []models.MerchantResponse

	// Execute Query
	rows, err := repository.db.Query(query, userId)

	// Check if there is error
	if err != nil {
		return listMerchant, err
	}

	// Loop Each Data
	for rows.Next() {
		// Create Single Merchant Variable
		var merchant models.MerchantResponse

		// Scan Row to Struct
		if err := rows.Scan(
			&merchant.Id,
			&merchant.UserId,
			&merchant.MerchantName,
			&merchant.CreatedAt,
			&merchant.UpdatedAt,
			&merchant.DeletedAt,
		); err != nil {
			return listMerchant, err
		}

		listMerchant = append(listMerchant, merchant)
	}

	// Check rows error
	if err := rows.Err(); err != nil {
		return listMerchant, err
	}

	// Return Success Scan
	return listMerchant, nil
}

// Get Merchant Detail by User Id & Merchant Id
func (repository *merchantRepository) GetMerchantDetailByUserIdAndMerchantId(userId int, merchantId int) (models.MerchantResponse, error) {
	// Prepare Query
	query := `SELECT * FROM merchants WHERE user_id=? AND id=? AND deleted_at IS NULL`

	// Prepare New Struct Instance
	var merchant models.MerchantResponse

	// Execute Query
	row := repository.db.QueryRow(query, userId, merchantId)

	// Scan Row to User Response Model
	if err := row.Scan(
		&merchant.Id,
		&merchant.UserId,
		&merchant.MerchantName,
		&merchant.CreatedAt,
		&merchant.UpdatedAt,
		&merchant.DeletedAt,
	); err != nil {
		return models.MerchantResponse{}, err
	}

	// Return User Response Model
	return merchant, nil
}
