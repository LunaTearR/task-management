package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title         string `gorm:"not null"`
	Description   string
	Priority      int
	CriticalLevel string
	Status        string
	ProjectID     uint
	Project       Project `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DeadlineAt    *time.Time
	StartedAt     *time.Time
	CompletedAt   *time.Time
	Assignees     []TaskUser `gorm:"foreignKey:TaskID"`
}

type TaskUser struct {
	TaskID uint `gorm:"primaryKey;autoIncrement:false"`
	UserID uint `gorm:"primaryKey;autoIncrement:false"`
}
 