package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string `gorm:"uniqueIndex;not null"`
	Password  string `gorm:"not null"`
	Email     string `gorm:"uniqueIndex;not null"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
}
