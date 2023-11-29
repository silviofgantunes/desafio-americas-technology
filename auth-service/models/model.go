// ./auth-service/models/model.go

package models

import (
	"time"
)

type Admin struct {
	ID        string    `gorm:"type:char(36);primary_key" json:"id,omitempty"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name,omitempty"`
	Email     string    `gorm:"type:varchar(255);not null" json:"email,omitempty"`
	Password  string    `gorm:"type:varchar(255);not null" json:"password,omitempty"`
	CreatedAt time.Time `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at,omitempty"`
}

type CreateAdminRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateAdminRequest struct {
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
