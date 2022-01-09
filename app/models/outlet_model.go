package models

import "time"

// Outlet ( Database )
type Outlet struct {
	UserId     int    `db:"user_id"`
	MerchantId int    `db:"merchant_id"`
	OutletName string `db:"outlet_name"`
}

// Outlet ( Request )
type OutletRequest struct {
	UserId     int    `json:"user_id" binding:"omitempty"`
	MerchantId int    `json:"merchant_id" binding:"omitempty"`
	OutletName string `json:"outlet_name" binding:"required"`
}

// Outlet ( Response )
type OutletResponse struct {
	Id         int        `json:"id" db:"id"`
	UserId     int        `json:"user_id" db:"user_id"`
	MerchantId int        `json:"merchant_id" db:"merchant_id"`
	OutletName string     `json:"outlet_name" db:"outlet_name"`
	CreatedAt  time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at" db:"deleted_at"`
}
