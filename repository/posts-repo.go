package repository

import (
	"github.com/leogsouza/go-rest-api/entity"
)

// PostRepository holds post collection operations
type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
	Delete(post *entity.Post) error
}
