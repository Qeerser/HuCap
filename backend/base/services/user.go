package services

import "HuCap/base/schemas"

type UserService interface {
	CreateUser(u *schemas.User) error
}