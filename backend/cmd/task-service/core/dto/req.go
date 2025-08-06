package dto

import "time"

type CreateTaskRequest struct {
	Title         string     `json:"title" binding:"required"`
	Description   string     `json:"description"`
	Priority      int        `json:"priority"`
	CriticalLevel string     `json:"critical_level"`
	Status        string     `json:"status"`
	StartedAt     *time.Time `json:"started_at"`
	CompletedAt   *time.Time `json:"completed_at"`
	DeadlineAt    *time.Time `json:"deadline_at"`
	ProjectID     uint       `json:"project_id" binding:"required"`
	UserIDs       []uint     `json:"user_ids"`
}

type CreateProjectRequest struct {
	Name        string     `json:"name" binding:"required"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	OwnerID     uint       `json:"owner_id" binding:"required"`
	StartedAt   *time.Time `json:"started_at"`
	FinishedAt  *time.Time `json:"finished_at"`
	DeadlineAt  *time.Time `json:"deadline_at"`
	UserIDs     []uint     `json:"user_ids"`
}
