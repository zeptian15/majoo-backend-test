package repositories

import (
	"database/sql"
	"fmt"
	"majoo-backend-test/app/models"
)

// Create New Outlet Repository Interface
type OutletRepository interface {
	CreateOutlet(models.Outlet) error
	UpdateOutlet(outlet models.Outlet, merchantId int) error
	DeleteOutlet(merchantId int) error
	GetOutletByOutletId(merchantId int) (models.OutletResponse, error)
	GetListOutletByUserId(userId int) ([]models.OutletResponse, error)
	GetListOutletByUserIdAndMerchantId(userId int, merchantId int) ([]models.OutletResponse, error)
	GetOutletDetailByUserIdAndOutletId(userId int, merchantId int) (models.OutletResponse, error)
	CheckIfOutletExistByOutletId(outletId int) (bool, error)
}

// Binding Database to Repository ( Constructor )
type outletRepository struct {
	db *sql.DB
}

// Create Function, to Create New Outlet Repository Instance
func NewOutletRepository(db *sql.DB) *outletRepository {
	return &outletRepository{db}
}

///
/// Implementing Outlet Repository Interface
///

// Create Outlet
func (repository *outletRepository) CreateOutlet(outlet models.Outlet) error {
	// Prepare Query
	query := `INSERT INTO outlets ( user_id, merchant_id, outlet_name ) VALUES ( ?, ?, ? )`

	// Execute Query
	_, err := repository.db.Exec(query, outlet.UserId, outlet.MerchantId, outlet.OutletName)

	// Return err
	return err
}

// Update Outlet
func (repository *outletRepository) UpdateOutlet(outlet models.Outlet, outletId int) error {
	// Prepare Query
	query := `UPDATE outlets SET outlet_name=?, merchant_id=? WHERE id=?`

	fmt.Println(outlet.OutletName, outlet)

	// Execute Query
	_, err := repository.db.Exec(query, outlet.OutletName, outlet.MerchantId, outletId)

	// Return err
	return err
}

// Delete Outlet ( Soft Delete )
func (repository *outletRepository) DeleteOutlet(outletId int) error {
	// Prepare Query
	query := `UPDATE outlets SET deleted_at=NOW() WHERE id=?`

	// Execute Query
	_, err := repository.db.Exec(query, outletId)

	// Return err
	return err
}

// Get Outlet by Outlet Id
func (repository *outletRepository) GetOutletByOutletId(outletId int) (models.OutletResponse, error) {
	// Prepare Query
	query := `SELECT * FROM outlets WHERE id=? AND deleted_at IS NULL`

	// Prepare New Struct Instance
	var outlet models.OutletResponse

	// Execute Query
	row := repository.db.QueryRow(query, outletId)

	// Scan Row to User Response Model
	if err := row.Scan(
		&outlet.Id,
		&outlet.UserId,
		&outlet.MerchantId,
		&outlet.OutletName,
		&outlet.CreatedAt,
		&outlet.UpdatedAt,
		&outlet.DeletedAt,
	); err != nil {
		return models.OutletResponse{}, err
	}

	// Return Outlet Response Model
	return outlet, nil
}

// Get List Outlet by User Id
func (repository *outletRepository) GetListOutletByUserId(userId int) ([]models.OutletResponse, error) {
	// Prepare Query
	query := `SELECT * FROM outlets WHERE user_id=? AND deleted_at IS NULL`

	// Create List Outlet Response Variable
	var listOutlet []models.OutletResponse

	// Execute Query
	rows, err := repository.db.Query(query, userId)

	// Check if there is error
	if err != nil {
		return listOutlet, err
	}

	// Loop Each Data
	for rows.Next() {
		// Create Single Outlet Variable
		var outlet models.OutletResponse

		// Scan Row to Struct
		if err := rows.Scan(
			&outlet.Id,
			&outlet.UserId,
			&outlet.MerchantId,
			&outlet.OutletName,
			&outlet.CreatedAt,
			&outlet.UpdatedAt,
			&outlet.DeletedAt,
		); err != nil {
			return listOutlet, err
		}

		listOutlet = append(listOutlet, outlet)
	}

	// Check rows error
	if err := rows.Err(); err != nil {
		return listOutlet, err
	}

	// Return Success Scan
	return listOutlet, nil
}

// Get List Outlet by User Id
func (repository *outletRepository) GetListOutletByUserIdAndMerchantId(userId int, merchantId int) ([]models.OutletResponse, error) {
	// Prepare Query
	query := `SELECT * FROM outlets WHERE user_id=? AND merchant_id=? AND deleted_at IS NULL`

	// Create List Outlet Response Variable
	var listOutlet []models.OutletResponse

	// Execute Query
	rows, err := repository.db.Query(query, userId, merchantId)

	// Check if there is error
	if err != nil {
		return listOutlet, err
	}

	// Loop Each Data
	for rows.Next() {
		// Create Single Outlet Variable
		var outlet models.OutletResponse

		// Scan Row to Struct
		if err := rows.Scan(
			&outlet.Id,
			&outlet.UserId,
			&outlet.MerchantId,
			&outlet.OutletName,
			&outlet.CreatedAt,
			&outlet.UpdatedAt,
			&outlet.DeletedAt,
		); err != nil {
			return listOutlet, err
		}

		listOutlet = append(listOutlet, outlet)
	}

	// Check rows error
	if err := rows.Err(); err != nil {
		return listOutlet, err
	}

	// Return Success Scan
	return listOutlet, nil
}

// Get Outlet Detail by User Id & Outlet Id
func (repository *outletRepository) GetOutletDetailByUserIdAndOutletId(userId int, outletId int) (models.OutletResponse, error) {
	// Prepare Query
	query := `SELECT * FROM outlets WHERE user_id=? AND id=? AND deleted_at IS NULL`

	// Prepare New Struct Instance
	var outlet models.OutletResponse

	// Execute Query
	row := repository.db.QueryRow(query, userId, outletId)

	// Scan Row to Outlet Response Model
	if err := row.Scan(
		&outlet.Id,
		&outlet.UserId,
		&outlet.MerchantId,
		&outlet.OutletName,
		&outlet.CreatedAt,
		&outlet.UpdatedAt,
		&outlet.DeletedAt,
	); err != nil {
		return models.OutletResponse{}, err
	}

	// Return Outlet Response Model
	return outlet, nil
}

// Check if Outlet Exists by Outlet Id
func (repository *outletRepository) CheckIfOutletExistByOutletId(outletId int) (bool, error) {
	// Prepare Query
	query := `SELECT id FROM outlets WHERE id=? AND deleted_at IS NULL`

	// Execute Query
	err := repository.db.QueryRow(query, outletId).Scan(&outletId)

	// Check if there is error
	if err != nil {
		// If Error is not NoRows
		if err != sql.ErrNoRows {
			// Return Query Error
			return false, err
		}
		// If error is NoRows = Outlet does not exists
		return false, nil
	}

	// Return Outlet Exist
	return true, nil
}
