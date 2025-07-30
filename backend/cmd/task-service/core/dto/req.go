package dto

import "time"

type ReqTask struct {
	Title         string     `json:"title" binding:"required"`
	Description   string     `json:"description"`
	Status        string     `json:"status" binding:"required"`
	Priority      int        `json:"priority"`  // เพิ่ม Priority
	CriticalLevel string     `json:"critical_level"` // เพิ่ม CriticalLevel
	UserID        uint       `json:"user_id" binding:"required"` // เพิ่ม UserID สำหรับเชื่อมโยงกับ User
	DeadlineAt    *time.Time `json:"deadline_at"` // เพิ่ม DeadlineAt
	StartedAt     *time.Time `json:"started_at"`  // เพิ่ม StartedAt
	CompletedAt   *time.Time `json:"completed_at"` // เพิ่ม CompletedAt
}