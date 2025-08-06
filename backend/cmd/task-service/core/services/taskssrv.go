package services

import (
	"fmt"
	"task-management-task-service/core/dto"
	"task-management-task-service/core/interfaces"
)

type TaskService struct {
	repo interfaces.TaskRepository
}

func NewTaskService(repo interfaces.TaskRepository) interfaces.TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(req dto.CreateTaskRequest) error {

	err := s.repo.CreateTask(req)
	if err != nil {
		return fmt.Errorf("failed to create task: %w", err)
	}
	return nil
}

func (s *TaskService) GetTasks() ([]dto.Task, error) {
	tasks, err := s.repo.GetTasks()
	if err != nil {
		return nil, fmt.Errorf("failed to get tasks: %w", err)
	}
	return tasks, nil
}
