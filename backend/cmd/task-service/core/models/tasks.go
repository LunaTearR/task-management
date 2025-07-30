package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title         string
	Description   string
	Priority      int
	CriticalLevel string
	Status        string

	DeadlineAt  *time.Time `gorm:"default:null"`
	StartedAt   *time.Time `gorm:"default:null"`
	CompletedAt *time.Time `gorm:"default:null"`
}
