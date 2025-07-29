package models

import (
	"gorm.io/gorm"
)

const (
	RoleAdmin string = "admin"
	RoleUser  string = "user"
)

type User struct {
	gorm.Model
	Name         string `gorm:"not null" json:"name"`
	Email        string `gorm:"type:citext;not null;unique" json:"email"`
	Role         string `gorm:"type:varchar(20);not null;default:'user'"`
	PasswordHash string `gorm:"not null" json:"-"`
}
