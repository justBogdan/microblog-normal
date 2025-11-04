package repositories

import (
	"errors"
	"microblog-normal/internal/models"
	"sync"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type UserRepository interface {
	CreateUser(user *models.User) error
	FindByNick(nickname string) (models.User, error)
}

type UserRepo struct {
	mutex sync.RWMutex
	users map[string]*models.User
}

func NewUserRepository() *UserRepo {
	return &UserRepo{
		users: make(map[string]*models.User),
	}
}

func (repo *UserRepo) CreateUser(user *models.User) error {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()
	repo.users[user.Nickname] = user
	return nil
}

func (repo *UserRepo) FindByNick(nickname string) (models.User, error) {
	repo.mutex.RLock()
	defer repo.mutex.RUnlock()
	if user, exists := repo.users[nickname]; exists {
		return *user, nil
	}
	return models.User{}, ErrUserNotFound
}
