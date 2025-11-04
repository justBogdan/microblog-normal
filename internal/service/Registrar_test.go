package service

import (
	"microblog-normal/internal/models"
	"microblog-normal/internal/repositories"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type mocUserRepository struct {
	mock.Mock
}

func (mr *mocUserRepository) FindByNick(nickname string) (models.User, error) {
	args := mr.Called(nickname)
	return args.Get(0).(models.User), args.Error(1)
}

func (mr *mocUserRepository) CreateUser(user *models.User) error {
	args := mr.Called(user)
	return args.Error(0)
}
func TestRegistrar_Register_unique_user(t *testing.T) {
	mocRepo := new(mocUserRepository)
	mocRepo.On("CreateUser", mock.Anything).Return(nil)
	mocRepo.On("FindByNick", "Sonya").Return(models.User{}, repositories.ErrUserNotFound)
	registrar := NewRegistrar(mocRepo)
	user, err := registrar.Register("Sonya")
	require.NoError(t, err)
	require.NotZero(t, user.ID)
	require.Equal(t, "Sonya", user.Nickname)
}

func TestRegistrar_Register_ExistUser(t *testing.T) {
	mocRepo := new(mocUserRepository)
	mocRepo.On("CreateUser", mock.Anything).Return(nil)
	mocRepo.On("FindByNick", "Sonya").Return(models.User{"Sonya", 0}, nil)
	registrar := NewRegistrar(mocRepo)
	user, err := registrar.Register("Sonya")
	require.Error(t, err)
	require.Equal(t, "Sonya", user.Nickname)
}
