package repository

import (
	"github.com/leogsouza/go-rest-api/entity"
	"github.com/stretchr/testify/mock"
)

// MockRepository represents a mocked repository
type MockRepository struct {
	mock.Mock
}

// Save mocks the save post method
func (mock *MockRepository) Save(post *entity.Post) (*entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.Post), args.Error(1)
}

// FindAll mocks the findAll post method
func (mock *MockRepository) FindAll() ([]entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.Post), args.Error(1)
}
