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

type service struct {
	repo repository.PostRepository
}

// NewPostService creates a new PostService instance
func NewPostService(repo repository.PostRepository) PostService {
	return &service{
		repo,
	}
}

// Validate validates de post data
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

// Create performs a new post creation
func (s *service) Create(post *entity.Post) (*entity.Post, error) {
	post.ID = rand.Int63()
	return s.repo.Save(post)
}

// FindAll returns all posts
func (s *service) FindAll() ([]entity.Post, error) {
	return s.repo.FindAll()
}
