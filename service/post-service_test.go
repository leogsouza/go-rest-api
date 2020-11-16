package service

import (
	"testing"

	"github.com/leogsouza/go-rest-api/entity"
	"github.com/leogsouza/go-rest-api/repository"
	"github.com/stretchr/testify/assert"
)

func TestValidateEmptyPost(t *testing.T) {
	testService := NewPostService(nil)

	err := testService.Validate(nil)
	assert.NotNil(t, err)

	assert.Equal(t, "The post is empty", err.Error())

}

func TestValidateEmptyPostTitle(t *testing.T) {
	post := &entity.Post{ID: 1, Title: "", Text: "A"}

	testService := NewPostService(nil)
	err := testService.Validate(post)

	assert.NotNil(t, err)

	assert.Equal(t, "The post title is empty", err.Error())

}

func TestFindAll(t *testing.T) {
	mockRepo := new(repository.MockRepository)

	var identifier int64 = 1

	post := entity.Post{ID: identifier, Title: "A", Text: "B"}

	// Setup expectations
	mockRepo.On("FindAll").Return([]entity.Post{post}, nil)

	testService := NewPostService(mockRepo)

	result, _ := testService.FindAll()

	// Mock Assert: Behavioral
	mockRepo.AssertExpectations(t)

	// Data Assertion
	assert.Equal(t, identifier, result[0].ID)
	assert.Equal(t, "A", result[0].Title)
	assert.Equal(t, "B", result[0].Text)
}

func TestCreate(t *testing.T) {
	var identifier int64 = 1

	mockRepo := new(repository.MockRepository)
	post := entity.Post{ID: identifier, Title: "A", Text: "B"}

	// Setup expectations
	mockRepo.On("Save").Return(&post, nil)

	testService := NewPostService(mockRepo)
	result, err := testService.Create(&post)

	// Mock Assert: Behavioral
	mockRepo.AssertExpectations(t)

	// Data Assertion
	assert.NotNil(t, result.ID)
	assert.Equal(t, "A", result.Title)
	assert.Equal(t, "B", result.Text)
	assert.Nil(t, err)

}
