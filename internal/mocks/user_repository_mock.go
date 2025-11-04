package mocks

import (
	"microblog-normal/internal/models"

	"github.com/stretchr/testify/mock"
)

type MocUserRepository struct {
	mock.Mock
}

func (mr *MocUserRepository) FindByNick(nickname string) (models.User, error) {
	args := mr.Called(nickname)
	return args.Get(0).(models.User), args.Error(1)
}

func (mr *MocUserRepository) CreateUser(user *models.User) error {
	args := mr.Called(user)
	return args.Error(0)
}
