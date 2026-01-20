package services

import (
	"fmt"

	"github.com/Yogi-1996/notes-backend/internal/models"
	"github.com/Yogi-1996/notes-backend/internal/repository"
	"github.com/Yogi-1996/notes-backend/pkg/hash"
	"github.com/Yogi-1996/notes-backend/pkg/jwt"
)

type UserServiceInterface interface {
	AddUser(email, password string) (models.User, error)
	VerifyUser(email, password string) (string, error)
}

type UserService struct {
	repo repository.UserRepositoryInterface
}

func NewUserService(u repository.UserRepositoryInterface) *UserService {
	return &UserService{
		repo: u,
	}
}

func (u *UserService) AddUser(email, password string) (models.User, error) {
	_, err := u.repo.GetUserByEmail(email)

	if err == nil {
		return models.User{}, fmt.Errorf("Emailid already registered: %w", err)
	}

	hashpassword, err := hash.GenerateHash(password)
	if err != nil {
		return models.User{}, fmt.Errorf("password cannot be hashed: %w", err)
	}

	user := models.User{
		Email:    email,
		Password: hashpassword,
	}
	err = u.repo.AddUser(&user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (u *UserService) VerifyUser(email, password string) (string, error) {
	user, err := u.repo.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	hashpassword := user.Password

	passwordcheck := hash.VerifyPassword(password, hashpassword)
	if !passwordcheck {
		return "", fmt.Errorf("Password Doesnt match")
	}

	token, err := jwt.GenerateToken(user.ID)

	if err != nil {
		return "", err
	}

	return token, nil
}
