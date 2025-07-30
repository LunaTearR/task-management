package repositories

import (
	"task-management-task-service/core/dto"
	"task-management-task-service/core/interfaces"
	"task-management-task-service/core/models"

	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) interfaces.TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) CreateTask(req dto.ReqTask) error {

	task := models.Task{
		Title:         req.Title,
		Description:   req.Description,
		Status:        req.Status,
		Priority:      req.Priority,
		CriticalLevel: req.CriticalLevel,
		DeadlineAt:    req.DeadlineAt,
		StartedAt:     req.StartedAt,
		CompletedAt:   req.CompletedAt,
	}

	if err := r.db.Create(&task).Error; err != nil {
		return err
	}

	return nil
}

func (r *TaskRepository) GetTasks() ([]dto.Task, error) {
	var tasks []dto.Task
	err := r.db.Model(&models.Task{}).Select("id, title, description, status, priority, critical_level, user_id, deadline_at, started_at, completed_at").Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
