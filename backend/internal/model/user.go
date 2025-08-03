package model

import (
	"time"
)

// User represents the user entity
type User struct {
	ID        int64     `json:"id" db:"id"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"-" db:"password"` // Không hiển thị password trong JSON
	FirstName string    `json:"firstName,omitempty" db:"first_name"`
	LastName  string    `json:"lastName,omitempty" db:"last_name"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}

// UserSignup represents data required for user registration
type UserSignup struct {
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=6"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
}

// UserLogin represents data required for user login
type UserLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// UserResponse represents the user data to be returned in API responses
type UserResponse struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"firstName,omitempty"`
	LastName  string    `json:"lastName,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
}

// TokenResponse represents the token data to be returned after authentication
type TokenResponse struct {
	AccessToken string       `json:"accessToken"`
	User        UserResponse `json:"user"`
}
