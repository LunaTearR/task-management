package services

import (
	"fmt"
	"task-management-task-service/core/dto"
	"task-management-task-service/core/interfaces"
	"task-management-task-service/pkg/userclient"
)

type ProjectService struct {
	repo       interfaces.ProjectRepository
	userClient *userclient.UserClient
}

func NewProjectService(repo interfaces.ProjectRepository, userClient *userclient.UserClient) interfaces.ProjectService {
	return &ProjectService{repo: repo, userClient: userClient}
}

func (s *ProjectService) CreateProject(req dto.CreateProjectRequest) error {
	err := s.repo.CreateProject(req)
	if err != nil {
		return fmt.Errorf("failed to create project: %w", err)
	}
	return nil
}

func (s *ProjectService) GetProjects() ([]dto.ProjectResponse, error) {
	// 1. ดึง Projects
	projects, err := s.repo.GetProjects()
	if err != nil {
		return nil, err
	}

	// 2. รวม unique user IDs
	userIDSet := make(map[uint]struct{})
	for _, p := range projects {
		userIDSet[p.OwnerID] = struct{}{}
		for _, uid := range p.Members {
			userIDSet[uid] = struct{}{}
		}
	}

	// แปลง map เป็น slice
	userIDs := make([]uint, 0, len(userIDSet))
	for id := range userIDSet {
		userIDs = append(userIDs, id)
	}

	// 3. ดึง user จาก user-service
	users, err := s.userClient.GetUsersByIDs(userIDs)
	if err != nil {
		return nil, err
	}

	// 4. สร้าง map สำหรับ lookup
	userMap := make(map[uint]dto.UserInfo)
	for _, u := range users {
		userMap[u.ID] = dto.UserInfo{
			ID:        u.ID,
			Username:  u.Username,
			Email:     u.Email,
			FirstName: u.FirstName,
			LastName:  u.LastName,
		}
	}

	// 5. ประกอบข้อมูล ProjectResponse
	result := make([]dto.ProjectResponse, 0, len(projects))
	for _, p := range projects {
		var members []dto.UserInfo
		for _, uid := range p.Members {
			if u, ok := userMap[uid]; ok {
				members = append(members, u)
			}
		}
		result = append(result, dto.ProjectResponse{
			ID:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			Status:      p.Status,
			StartedAt:   p.StartedAt,
			FinishedAt:  p.FinishedAt,
			DeadlineAt:  p.DeadlineAt,
			Owner:       userMap[p.OwnerID],
			Members:     members,
		})
	}

	return result, nil
}
