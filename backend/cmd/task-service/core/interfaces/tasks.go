package interfaces

import (
	"task-management-task-service/core/dto"
)

type TaskRepository interface {
	CreateTask(req dto.CreateTaskRequest) error
	GetTasks() ([]dto.Task, error)
}

type TaskService interface {
	CreateTask(req dto.CreateTaskRequest) error
	GetTasks() ([]dto.Task, error)
}
