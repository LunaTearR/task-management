package repositories

import (
	"task-management-task-service/core/dto"
	"task-management-task-service/core/interfaces"
	"task-management-task-service/core/models"

	"gorm.io/gorm"
)

type ProjectRepository struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) interfaces.ProjectRepository {
	return &ProjectRepository{db: db}
}

func (r *ProjectRepository) CreateProject(req dto.CreateProjectRequest) error {
	project := models.Project{
		Name:        req.Name,
		Description: req.Description,
		Status:      req.Status,
		OwnerID:     req.OwnerID,
		StartedAt:   req.StartedAt,
		FinishedAt:  req.FinishedAt,
		DeadlineAt:  req.DeadlineAt,
	}

	// Step 1: Save project first to get project.ID
	if err := r.db.Create(&project).Error; err != nil {
		return err
	}

	// Step 2: Then assign Members with valid project.ID
	if len(req.UserIDs) > 0 {
		var members []models.ProjectMember
		for _, userID := range req.UserIDs {
			members = append(members, models.ProjectMember{
				ProjectID: project.ID,
				UserID:    userID,
			})
		}
		if err := r.db.Create(members).Error; err != nil {
			return err
		}
	}

	return nil
}

func (r *ProjectRepository) GetProjects() ([]dto.Project, error) {
	var projects []dto.Project
	err := r.db.Model(&models.Project{}).
		Select("id, name, description, status, owner_id, started_at, finished_at, deadline_at").
		Find(&projects).Error
	if err != nil {
		return nil, err
	}

	for i, project := range projects {
		var members []models.ProjectMember
		if err := r.db.Where("project_id = ?", project.ID).Find(&members).Error; err != nil {
			return nil, err
		}
		memberIDs := make([]uint, len(members))
		for j, m := range members {
			memberIDs[j] = m.UserID
		}
		projects[i].Members = memberIDs
	}

	if len(projects) == 0 {
		return nil, nil
	}

	return projects, nil
}
