package interfaces

import (
	"task-management-task-service/core/dto"
)

type ProjectRepository interface {
	CreateProject(req dto.CreateProjectRequest) error
	GetProjects() ([]dto.Project, error)
}

type ProjectService interface {
	CreateProject(req dto.CreateProjectRequest) error
	GetProjects() ([]dto.ProjectResponse, error)
}
