package services

import (
	"HuCap/base/schemas"
)

type UserService interface {
	CreateUser(u *schemas.User) error
	LoginUser(user *schemas.User) bool
}