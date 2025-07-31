package services

import (
	"fmt"
	"task-management-user-service/core/dto"
	"task-management-user-service/core/interfaces"
)

type UserService struct {
	repo interfaces.UserRepository
}

func NewUserService(repo interfaces.UserRepository) interfaces.UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(req dto.ReqUser) error {

	err := s.repo.CreateUser(req)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func (s *UserService) GetUsers() ([]dto.User, error) {
	users, err := s.repo.GetUsers()
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}
	return users, nil
}

func (s *UserService) GetUserByID(id uint) (*dto.User, error) {
	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by ID: %w", err)
	}
	return user, nil
}

func (s *UserService) UpdateUser(user *dto.ReqUser, id int) error {
	if err := s.repo.UpdateUser(user, id); err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}

func (s *UserService) DeleteUser(id uint) error {

	if err := s.repo.DeleteUser(id); err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}
