package services

import (
	"HuCap/base/models"
	"HuCap/base/repositories"
	"HuCap/base/schemas"
	"HuCap/utils"

	"github.com/google/uuid"
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

func (s *userService) LoginUser(user *schemas.User) bool {
	dsn := s.userRepo.LoginUser(user.Email)
	user.ID , _ = uuid.Parse(dsn.ID)
	user.FullName = dsn.FullName
	user.Username = dsn.Username
	user.Password = ""
	return utils.CheckPasswordHash(user.Password ,dsn.Password)
}