package service

import (
	"errors"
	"math/rand"
	"microblog-normal/internal/models"
	"microblog-normal/internal/repositories"
)

type Registrar interface {
	Register(nickname string) (models.User, error)
}

type RegisterService struct {
	userRepo repositories.UserRepository
}

func NewRegistrar(userRepo repositories.UserRepository) *RegisterService {
	return &RegisterService{userRepo: userRepo}
}

var (
	ErrUserAlreadyExists = errors.New("user already exists")
)

func (rs *RegisterService) Register(nickname string) (models.User, error) {
	if user, userExist := rs.userRepo.FindByNick(nickname); userExist == nil {
		return user, ErrUserAlreadyExists // вот здесь вопрос что лучше использовать bool flag?
	}
	id := rand.Int31()
	registerUser := models.User{Nickname: nickname, ID: id}
	if err := rs.userRepo.CreateUser(&registerUser); err != nil {
		return models.User{}, err
	}
	return registerUser, nil
}
