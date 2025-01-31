package services

import "HuCap/base/repositories"

type service struct {
	userService UserService
}

func NewService(repo repositories.Repository) Service{
	return &service{
		userService: NewUserService(repo.User()),
	}
}

func (s *service) User() UserService {
	return s.userService
}