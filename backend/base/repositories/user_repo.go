package repositories

import (
	"HuCap/base/models"
	"log"

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

func (r *userRepository) LoginUser(email string) *models.User {
	var user models.User
	result := r.DB.Where("email = ?", email).Find(&user)
	if result.Error != nil {
	  log.Fatalf("Error finding book: %v", result.Error)
	}

	return &user
}