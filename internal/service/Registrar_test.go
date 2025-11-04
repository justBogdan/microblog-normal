package service

import (
	"microblog-normal/internal/mocks"
	"microblog-normal/internal/models"
	"microblog-normal/internal/repositories"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestRegistrar_Register_unique_user(t *testing.T) {
	mocRepo := new(mocks.MocUserRepository)
	mocRepo.On("CreateUser", mock.Anything).Return(nil)
	mocRepo.On("FindByNick", "Sonya").Return(models.User{}, repositories.ErrUserNotFound)
	registrar := NewRegistrar(mocRepo)
	user, err := registrar.Register("Sonya")
	require.NoError(t, err)
	require.NotZero(t, user.ID)
	require.Equal(t, "Sonya", user.Nickname)
}

func TestRegistrar_Register_ExistUser(t *testing.T) {
	mocRepo := new(mocks.MocUserRepository)
	mocRepo.On("CreateUser", mock.Anything).Return(nil)
	mocRepo.On("FindByNick", "Sonya").Return(models.User{"Sonya", 0}, nil)
	registrar := NewRegistrar(mocRepo)
	user, err := registrar.Register("Sonya")
	require.Error(t, err)
	require.Equal(t, "Sonya", user.Nickname)
}
