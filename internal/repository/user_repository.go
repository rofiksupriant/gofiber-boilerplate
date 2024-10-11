package repository

import (
	"boilerplate/internal/entiity"
	"gorm.io/gorm"
)

type UserRepository struct {
	Repository[entiity.User]
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) Create(db *gorm.DB, user *entiity.User) error {
	err := db.Create(user).Error
	if err != nil {
		return err
	}

	return nil
}
