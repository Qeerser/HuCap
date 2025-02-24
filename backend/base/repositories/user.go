package repositories

import "HuCap/base/models"

type UserRepository interface {
	CreateUser(s *models.User) error
	LoginUser(email string) *models.User
}