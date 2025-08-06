package interfaces

import (
	"task-management-user-service/core/dto"
)

type UserRepository interface {
	CreateUser(req dto.ReqUser) error
	GetUsers() ([]dto.User, error)
	GetUserByID(id []uint) ([]dto.User, error)
	UpdateUser(user *dto.ReqUser, id int) error
	DeleteUser(id uint) error
}

type UserService interface {
	CreateUser(req dto.ReqUser) error
	GetUsers() ([]dto.User, error)
	GetUserByID(id []uint) ([]dto.User, error)
	UpdateUser(user *dto.ReqUser, id int) error
	DeleteUser(id uint) error
}
