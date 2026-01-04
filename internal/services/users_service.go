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
	hashpassword, err := hash.GenerateHash(password)
	if err != nil {
		return models.User{}, fmt.Errorf("password cannot be hashed: %w", err)
	}

	newuser, check := u.repo.AddUser(email, hashpassword)
	if !check {
		return models.User{}, fmt.Errorf("User Already exist")
	}

	return newuser, nil
}

func (u *UserService) VerifyUser(email, password string) (string, error) {
	user, check := u.repo.GetUser(email)
	if !check {
		return "", fmt.Errorf("user does not exist")
	}

	hashpassword := user.Password

	passwordcheck := hash.VerifyPassword(password, hashpassword)
	if !passwordcheck {
		return "", fmt.Errorf("Password Doesnt match")
	}

	token, err := jwt.GenerateToken(email)

	if err != nil {
		return "", err
	}

	return token, nil
}
