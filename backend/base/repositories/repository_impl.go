package repositories

import "gorm.io/gorm"

type repository struct {
	UserRepository UserRepository
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		UserRepository: NewUserRepository(db),
	}
}

func (r *repository) User() UserRepository{
	return r.UserRepository
}
