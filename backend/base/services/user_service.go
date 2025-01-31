package services

import (
	"HuCap/base/models"
	"HuCap/base/repositories"
	"HuCap/base/schemas"
	"HuCap/utils"
)

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) CreateUser(u *schemas.User) error{

	Hashed_password , _ := utils.HashPassword(u.Password)
	dnp := &models.User{      
		Email     : u.Email,
		Username  : u.Username,
		Password  : Hashed_password,
		FullName  : u.FullName,
	}

	return s.userRepo.CreateUser(dnp)
}