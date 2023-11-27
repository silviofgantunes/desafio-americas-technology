package models

import (
	"time"
)

type User struct {
	ID          string    `gorm:"type:char(36);primary_key" json:"id,omitempty"`
	Name        string    `gorm:"type:varchar(255);not null" json:"name,omitempty"`
	Email       string    `gorm:"type:varchar(255);not null" json:"email,omitempty"`
	PhoneNumber string    `gorm:"type:varchar(255);not null" json:"phone_number,omitempty"`
	CreatedAt   time.Time `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt   time.Time `gorm:"not null" json:"updated_at,omitempty"`
}

type CreateUserRequest struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
}

type UpdateUserRequest struct {
	Name        string `json:"name,omitempty"`
	Email       string `json:"email,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
}
