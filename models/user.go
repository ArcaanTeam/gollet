package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string `gorm:"not null" json:"name"`
	Email        string `gorm:"type:citext;not null;unique" json:"email"`
	PasswordHash string `gorm:"not null"`
}
