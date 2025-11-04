package repositories

import (
	"microblog-normal/internal/models"
	"sync"
)

type PostRepository interface {
	CreatePost(p *models.Post) error
	GetAllPosts() []models.Post
}

type PostRepo struct {
	mutex sync.RWMutex
	posts []*models.Post
}

func NewPostRepository() *PostRepo {
	return &PostRepo{
		posts: make([]*models.Post, 0, 50),
	}
}

func (repo *PostRepo) CreatePost(p *models.Post) error {
	// Cпросить у бадди нужна ли здесь проверка на nil указатель)))
	repo.mutex.Lock()
	defer repo.mutex.Unlock()
	repo.posts = append(repo.posts, p)
	return nil
}

func (repo *PostRepo) GetAllPosts() []models.Post {
	repo.mutex.RLock()
	defer repo.mutex.RUnlock()
	var copyPosts []models.Post
	for _, post := range repo.posts {
		copyPosts = append(copyPosts, *post) // протестить просто руками как это работает и меняется ли ориг
	}
	return copyPosts
}
