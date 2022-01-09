package repositories

import (
	"database/sql"
	"majoo-backend-test/app/models"
	"time"
)

// Create New Transaction Repository Interface
type TransactionRepository interface {
	GetListTransactionReportByUserId(userId int) ([]models.TransactionResponse, error)
	GetListTransactionReportDetailByUserId(userId int) ([]models.TransactionDetailResponse, error)
}

// Binding Database to Repository ( Constructor )
type transactionRepository struct {
	db *sql.DB
}

// Create Function, to Create New Transaction Repository Instance
func NewTransactionRepository(db *sql.DB) *transactionRepository {
	return &transactionRepository{db}
}

///
/// Implementing Transaction Repository Interface
///

// Get List Transaction by User Id
func (repository *transactionRepository) GetListTransactionReportByUserId(userId int) ([]models.TransactionResponse, error) {
	// Create Start Date
	startDate := time.Date(2021, 11, 1, 0, 0, 0, 0, time.UTC)

	// Create Container Variable
	var listTransaction []models.TransactionResponse

	// Loop 30 Days
	for index := 0; index < 30; index++ {
		// Check data From Database
		query := `SELECT t.id, m.user_id, m.merchant_name, SUM(t.bill_total) as omzet, t.created_at, t.updated_at, t.deleted_at
			FROM transactions t
			INNER JOIN merchants m
				ON t.merchant_id = m.id
			WHERE m.user_id=? AND t.created_at LIKE CONCAT(?, "%")
			GROUP BY t.merchant_id`

		// Execute Query
		rows, err := repository.db.Query(query, userId, startDate.Format("2006-01-02"))

		// Check if there is error
		if err != nil {
			return []models.TransactionResponse{}, err
		}

		// Create Temp Variable to Check is date is null ( 0001-01-01 )
		var lastDate time.Time

		// Loop Each Data
		for rows.Next() {
			// Create New Single Transaction Detail Response
			var transaction models.TransactionResponse

			// Scan Row to Transaction Detail Response Model
			if err := rows.Scan(
				&transaction.Id,
				&transaction.UserId,
				&transaction.MerchantName,
				&transaction.BillTotal,
				&transaction.CreatedAt,
				&transaction.UpdatedAt,
				&transaction.DeletedAt,
			); err != nil {
				return []models.TransactionResponse{}, err
			}

			// Append data to List
			listTransaction = append(listTransaction, transaction)
			lastDate = transaction.CreatedAt
		}

		if lastDate.Format("2006-01-02") == "0001-01-01" {
			// Append Empty Data
			var emptyTransaction models.TransactionResponse
			emptyTransaction.Id = 0
			emptyTransaction.UserId = 0
			emptyTransaction.MerchantName = ""
			emptyTransaction.BillTotal = 0
			emptyTransaction.CreatedAt = startDate
			emptyTransaction.UpdatedAt = startDate
			emptyTransaction.DeletedAt = nil

			// Append data to List
			listTransaction = append(listTransaction, emptyTransaction)
		}

		// Add 1 day to date
		startDate = startDate.AddDate(0, 0, 1)
	}

	// Return Success Scan
	return listTransaction, nil
}

// Get List Transaction by User Id
func (repository *transactionRepository) GetListTransactionReportDetailByUserId(userId int) ([]models.TransactionDetailResponse, error) {
	// Create Start Date
	startDate := time.Date(2021, 11, 1, 0, 0, 0, 0, time.UTC)

	// Create Container Variable
	var listTransaction []models.TransactionDetailResponse

	// Loop 30 Days
	for index := 0; index < 30; index++ {
		// Check data From Database
		query := `SELECT t.id, m.user_id, m.merchant_name, o.outlet_name, SUM(t.bill_total) as omzet, t.created_at, t.updated_at, t.deleted_at
			FROM transactions t
			INNER JOIN merchants m
				ON t.merchant_id = m.id
			INNER JOIN outlets o
				ON t.outlet_id = o.id
			WHERE m.user_id=? AND t.created_at LIKE CONCAT(?, "%")
			GROUP BY t.outlet_id`

		// Execute Query
		rows, err := repository.db.Query(query, userId, startDate.Format("2006-01-02"))

		// Check if there is error
		if err != nil {
			return []models.TransactionDetailResponse{}, err
		}

		// Create Temp Variable to Check is date is null ( 0001-01-01 )
		var lastDate time.Time

		// Loop Each Data
		for rows.Next() {
			// Create New Single Transaction Detail Response
			var transaction models.TransactionDetailResponse

			// Scan Row to Transaction Detail Response Model
			if err := rows.Scan(
				&transaction.Id,
				&transaction.UserId,
				&transaction.MerchantName,
				&transaction.Outlet_Name,
				&transaction.BillTotal,
				&transaction.CreatedAt,
				&transaction.UpdatedAt,
				&transaction.DeletedAt,
			); err != nil {
				return []models.TransactionDetailResponse{}, err
			}

			// Append data to List
			listTransaction = append(listTransaction, transaction)
			lastDate = transaction.CreatedAt
		}

		if lastDate.Format("2006-01-02") == "0001-01-01" {
			// Append Empty Data
			var emptyTransaction models.TransactionDetailResponse
			emptyTransaction.Id = 0
			emptyTransaction.UserId = 0
			emptyTransaction.MerchantName = ""
			emptyTransaction.BillTotal = 0
			emptyTransaction.CreatedAt = startDate
			emptyTransaction.UpdatedAt = startDate
			emptyTransaction.DeletedAt = nil

			// Append data to List
			listTransaction = append(listTransaction, emptyTransaction)
		}

		// Add 1 day to date
		startDate = startDate.AddDate(0, 0, 1)
	}

	// Return Success Scan
	return listTransaction, nil
}
