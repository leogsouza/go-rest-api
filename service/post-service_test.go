package service

import (
	"testing"

	"github.com/leogsouza/go-rest-api/entity"
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
