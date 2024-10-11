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

func (r *UserRepository) FindByUsername(db *gorm.DB, username string) (*entiity.User, error) {
	user := new(entiity.User)
	result := db.Model(&entiity.User{Username: username}).Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
