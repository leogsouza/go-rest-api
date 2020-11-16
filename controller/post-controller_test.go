package controller

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/leogsouza/go-rest-api/entity"
	"github.com/leogsouza/go-rest-api/repository"
	"github.com/leogsouza/go-rest-api/service"
	"github.com/stretchr/testify/assert"
)

const (
	ID    int64  = 123
	TITLE string = "Title 1"
	TEXT  string = "Text 1"
)

var (
	postRepo       repository.PostRepository = repository.NewSQLiteRepository()
	postService    service.PostService       = service.NewPostService(postRepo)
	postController PostController            = NewPostController(postService)
)

func TestAddPost(t *testing.T) {
	// Create a new HTTP POST request
	jsonData := []byte(`{"title": "Title 1", "text": "Text 1"}`)
	req, _ := http.NewRequest("POST", "/posts", bytes.NewBuffer(jsonData))

	// Assing HTTP handler function (controller AddPost function)
	handler := http.HandlerFunc(postController.AddPost)

	// Record HTTP Response (httptest)
	response := httptest.NewRecorder()

	// Dispatch the HTTP request
	handler.ServeHTTP(response, req)

	// Add Assersions on the HTTP Status Code and the response
	status := response.Code
	if status != http.StatusOK {
		t.Errorf("Handler returned a wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Decode the HTTP response
	var post entity.Post
	json.NewDecoder(io.Reader(response.Body)).Decode(&post)

	// Assert HTTP Response
	assert.NotNil(t, post.ID)
	assert.Equal(t, TITLE, post.Title)
	assert.Equal(t, TEXT, post.Text)

	// clean up database
	cleanUp(&post)
}

func TestGetPosts(t *testing.T) {
	// Insert new post
	setup()

	// Create a GET HTTP request
	req, _ := http.NewRequest("GET", "/posts", nil)

	// Assing HTTP handler function (controller GetPosts function)
	handler := http.HandlerFunc(postController.GetPosts)

	// Record HTTP Response (httptest)
	response := httptest.NewRecorder()

	// Dispatch the HTTP request
	handler.ServeHTTP(response, req)

	// Add Assersions on the HTTP Status Code and the response
	status := response.Code
	if status != http.StatusOK {
		t.Errorf("Handler returned a wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Decode the HTTP response
	var posts []entity.Post
	json.NewDecoder(io.Reader(response.Body)).Decode(&posts)

	// Assert HTTP Response
	assert.NotNil(t, posts[0].ID)
	assert.Equal(t, TITLE, posts[0].Title)
	assert.Equal(t, TEXT, posts[0].Text)

	// clean up database
	cleanUp(&posts[0])
}

func cleanUp(post *entity.Post) {
	postRepo.Delete(post)
}

func setup() {
	var post entity.Post = entity.Post{
		ID:    ID,
		Title: TITLE,
		Text:  TEXT,
	}
	postRepo.Save(&post)
}
