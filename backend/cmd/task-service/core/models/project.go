package models

import (
	"time"

	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Status      string          `json:"status"`
	OwnerID     uint            `json:"owner_id"`
	StartedAt   *time.Time      `json:"started_at"`
	FinishedAt  *time.Time      `json:"finished_at"`
	DeadlineAt  *time.Time      `json:"deadline_at"`
	Members     []ProjectMember `json:"members" gorm:"foreignKey:ProjectID;references:ID"`
}

type ProjectMember struct {
	ProjectID uint `json:"project_id" gorm:"primaryKey;autoIncrement:false"`
	UserID    uint `json:"user_id" gorm:"primaryKey;autoIncrement:false"`
}
