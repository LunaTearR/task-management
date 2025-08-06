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

func (r *TaskRepository) CreateTask(req dto.CreateTaskRequest) error {
	task := models.Task{
		Title:         req.Title,
		Description:   req.Description,
		Priority:      req.Priority,
		CriticalLevel: req.CriticalLevel,
		Status:        req.Status,
		DeadlineAt:    req.DeadlineAt,
		StartedAt:     req.StartedAt,
		CompletedAt:   req.CompletedAt,
		ProjectID:     req.ProjectID,
	}

	for _, userID := range req.UserIDs {
		task.Assignees = append(task.Assignees, models.TaskUser{UserID: userID})
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
