package models

// Register ( Request )
type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login ( Request )
type LoginRequest struct {
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}
