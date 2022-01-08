package models

import "time"

// User ( Database )
type User struct {
	Name     string `db:"name"`
	UserName string `db:"user_name"`
	Password string `db:"password"`
}

// User ( Request )
type UserRequest struct {
	Name     string `json:"name"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

// User ( Response )
type UserResponse struct {
	Id        int        `json:"id" db:"id"`
	Name      string     `json:"name" db:"name"`
	UserName  string     `json:"user_name" db:"user_name"`
	Password  string     `json:"password" db:"password"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
}

// User ( JWT Claims )
type UserClaims struct {
	Id       int    `json:"id"`
	UserName string `json:"user_name"`
}
