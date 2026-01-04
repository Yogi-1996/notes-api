package repository

import (
	"sync"
	"time"

	"github.com/Yogi-1996/notes-backend/internal/models"
)

type UserRepositoryInterface interface {
	AddUser(email, password string) (models.User, bool)
	GetUser(email string) (models.User, bool)
}

type UserRepository struct {
	mu     sync.Mutex
	users  map[int]models.User
	nextID int
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: make(map[int]models.User),
		nextID: 1,
	}
}

func (u *UserRepository) AddUser(email, password string) (models.User, bool) {
	u.mu.Lock()
	defer u.mu.Unlock()

	for _, user := range u.users {
		if user.Email == email {
			return models.User{}, false
		}
	}

	newuser := models.User{
		ID:        u.nextID,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	u.users[u.nextID] = newuser
	u.nextID += 1

	return newuser, true
}

func (u *UserRepository) GetUser(email string) (models.User, bool) {
	u.mu.Lock()
	defer u.mu.Unlock()

	for _, user := range u.users {
		if user.Email == email {
			return user, true
		}
	}

	return models.User{}, false
}
