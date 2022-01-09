package models

import "time"

// Transaction ( Response )
type TransactionResponse struct {
	Id           int        `json:"id" db:"id"`
	UserId       int        `json:"user_id" db:"user_id"`
	MerchantName string     `json:"merchant_name" db:"merchant_name"`
	BillTotal    uint64     `json:"omzet" db:"omzet" `
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at" db:"deleted_at"`
}

// Transaction Detail ( Response )
type TransactionDetailResponse struct {
	Id           int        `json:"id" db:"id"`
	UserId       int        `json:"user_id" db:"user_id"`
	MerchantName string     `json:"merchant_name" db:"merchant_name"`
	Outlet_Name  string     `json:"outlet_name" db:"outlet_name"`
	BillTotal    uint64     `json:"omzet" db:"omzet" `
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at" db:"deleted_at"`
}
