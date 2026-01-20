package repository

import (
	"github.com/Yogi-1996/notes-backend/internal/models"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	AddUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
}

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (u *UserRepository) AddUser(user *models.User) error {
	return u.DB.Create(user).Error
}

func (u *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	err := u.DB.First(&user, "email= ?", email).Error
	if err != nil {
		return nil, err
	}

	return &user, nil

}
