package repositories

import (
	"HuCap/base/models"

	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository (db *gorm.DB) UserRepository {
	return &userRepository{
		DB : db,
	}
}

func (r *userRepository) CreateUser(s *models.User) error {
	return r.DB.Save(s).Error
}