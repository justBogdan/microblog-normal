package mocks

import (
	"microblog-normal/internal/models"

	"github.com/stretchr/testify/mock"
)

type PostRepositoryMock struct {
	mock.Mock
}

func (repo *PostRepositoryMock) CreatePost(p *models.Post) error {
	args := repo.Called(p)
	return args.Error(0)
}

func (repo *PostRepositoryMock) GetAllPosts() []models.Post {
	args := repo.Called()
	return args.Get(0).([]models.Post)
}
