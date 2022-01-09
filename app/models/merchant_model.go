package models

import "time"

// Merchant ( Database )
type Merchant struct {
	UserId       int    `db:"user_id"`
	MerchantName string `db:"merchant_name"`
}

// Merchant ( Request )
type MerchantRequest struct {
	MerchantName string `json:"merchant_name" binding:"required"`
}

// Merchant ( Response )
type MerchantResponse struct {
	Id           int        `json:"id" db:"id"`
	UserId       int        `json:"user_id" db:"user_id"`
	MerchantName string     `json:"merchant_name" db:"merchant_name"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at" db:"deleted_at"`
}
