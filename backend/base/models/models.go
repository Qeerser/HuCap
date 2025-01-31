package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        string `gorm:"primary_key;default:uuid_generate_v4()"`
	Email     string    `gorm:"unique;not null"`
	Username  string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`
	FullName  string    `gorm:"not null"`
}