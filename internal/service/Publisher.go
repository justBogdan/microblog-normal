package service

import (
	"math/rand"
	"microblog-normal/internal/models"
	"microblog-normal/internal/repositories"
)

type Publisher interface {
	Publish(author string, text string) (models.Post, error)
	GetFeed() []models.Post
} // спросить у бадди верно ли что данный интерфейс вообще нужно объявить в ручке а здесь просто писать его имплементацию

type PostService struct {
	postRepo repositories.PostRepository
	userRepo repositories.UserRepository
}

func NewPublisher(postRepo repositories.PostRepository, userRepo repositories.UserRepository) *PostService {
	return &PostService{
		postRepo: postRepo,
		userRepo: userRepo,
	}
}

func (service *PostService) Publish(author string, text string) (models.Post, error) {
	if _, err := service.userRepo.FindByNick(author); err != nil {
		return models.Post{}, err
	}
	generateId := rand.Int31()
	newPost := &models.Post{
		Author: author,
		Text:   text,
		PostID: generateId, // вот здесь нет проверки на повторяющейся айди по началу ок, потом вводим проверку или влаг nextid
		Likes:  make([]models.Like, 0),
	}
	if err := service.postRepo.CreatePost(newPost); err != nil {
		return models.Post{}, err
	}
	return *newPost, nil
}

func (service *PostService) GetFeed() []models.Post {
	return service.postRepo.GetAllPosts() // можно сделать проверку на длину ленты)) if len(feed) == 0 return err
}
