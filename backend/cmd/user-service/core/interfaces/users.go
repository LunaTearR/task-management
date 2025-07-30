package interfaces

import (
	"task-management-user-service/core/dto"
)

type UserRepository interface {
	CreateUser(req dto.ReqUser) error
	GetUsers() ([]dto.User, error)
	DeleteUser(id uint) error
}

type UserService interface {
	CreateUser(req dto.ReqUser) error
	GetUsers() ([]dto.User, error)
	DeleteUser(id uint) error
}
