package repositories

import (
	"task-management-user-service/core/dto"
	"task-management-user-service/core/interfaces"
	"task-management-user-service/core/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) interfaces.UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(req dto.ReqUser) error {

	user := models.User{
		Username:  req.Username,
		Password:  req.Password,
		Email:     req.Email,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}
	if err := r.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetUsers() (result []dto.User, err error) {

	err = r.db.Model(&models.User{}).Select("id, username, email, first_name, last_name").Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *UserRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) UpdateUser(user *models.User) error {
	if err := r.db.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) DeleteUser(id uint) error {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return err
	}
	if err := r.db.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}
