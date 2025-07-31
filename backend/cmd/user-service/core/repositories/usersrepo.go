package repositories

import (
	"task-management-user-service/core/dto"
	"task-management-user-service/core/interfaces"
	"task-management-user-service/core/models"
	"task-management-user-service/utils"

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
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	if err := r.db.Where("email = ?", user.Email).First(&models.User{}).Error; err == nil {
		return utils.HandleResponse(nil, nil, "Email already exists", nil)
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

func (r *UserRepository) GetUserByID(id uint) (*dto.User, error) {
	var users *models.User
	if err := r.db.First(&users, id).Error; err != nil {
		return nil, err
	}

	user := &dto.User{
		Username:  users.Username,
		Email:     users.Email,
		FirstName: users.FirstName,
		LastName:  users.LastName,
	}
	return user, nil
}

func (r *UserRepository) UpdateUser(user *dto.ReqUser, id int) error {
	existingUser := &models.User{}

	if err := r.db.First(existingUser, id).Error; err != nil {
		return err
	}

	if user.Username != "" {
		existingUser.Username = user.Username
	}
	if user.Email != "" {
		existingUser.Email = user.Email
	}
	if user.FirstName != "" {
		existingUser.FirstName = user.FirstName
	}
	if user.LastName != "" {
		existingUser.LastName = user.LastName
	}

	if user.Password != "" {
		hashedPassword, err := utils.HashPassword(user.Password)
		if err != nil {
			return err
		}
		existingUser.Password = hashedPassword
	}

	if err := r.db.Save(existingUser).Error; err != nil {
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
