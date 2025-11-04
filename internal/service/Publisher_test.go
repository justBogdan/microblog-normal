package service

import (
	"microblog-normal/internal/mocks"
	"microblog-normal/internal/models"
	"microblog-normal/internal/repositories"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_Publisher_Publish_Success(t *testing.T) {
	mocUserRepo := new(mocks.MocUserRepository)
	mocPostRepo := new(mocks.PostRepositoryMock)
	mocPostRepo.On("CreatePost", mock.Anything).Return(nil)
	mocUserRepo.On("FindByNick", "Sonya").Return(models.User{"Sonya", 100}, nil)

	service := NewPublisher(mocPostRepo, mocUserRepo)

	result, err := service.Publish("Sonya", "Привет!")
	require.Nil(t, err)
	require.Equal(t, result.Text, "Привет!")
	require.Equal(t, result.Author, "Sonya")
	require.NotZero(t, result.PostID)
	require.Equal(t, result.Likes, []models.Like{})
}

func Test_Publisher_Publish_UserNotFount(t *testing.T) {
	mocUserRepo := new(mocks.MocUserRepository)
	mocPostRepo := new(mocks.PostRepositoryMock)
	mocPostRepo.On("CreatePost", mock.Anything).Return(nil)
	mocUserRepo.On("FindByNick", "Sonya").Return(models.User{}, repositories.ErrUserNotFound)

	service := NewPublisher(mocPostRepo, mocUserRepo)

	result, err := service.Publish("Sonya", "Привет!")
	require.Error(t, err)
	require.Equal(t, models.Post{}, result)
}
