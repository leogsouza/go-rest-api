package service

import (
	"errors"
	"math/rand"

	"github.com/leogsouza/go-rest-api/entity"
	"github.com/leogsouza/go-rest-api/repository"
)

type PostService interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

var (
	repo repository.PostRepository = repository.NewFirestoreRepository()
)

type service struct{}

func NewPostService() PostService {
	return &service{}
}

func (s *service) Validate(post *entity.Post) error {
	if post == nil {
		err := errors.New("The post is empty")
		return err
	}

	if post.Title == "" {
		err := errors.New("The post title is empty")
		return err
	}
	return nil
}

func (s *service) Create(post *entity.Post) (*entity.Post, error) {
	post.ID = rand.Int63()
	return repo.Save(post)
}

func (s *service) FindAll() ([]entity.Post, error) {
	return repo.FindAll()
}