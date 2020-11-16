package repository

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/leogsouza/go-rest-api/entity"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type repo struct{}

// NewFirestoreRepository returns a new instance which implements PostRepository
func NewFirestoreRepository() PostRepository {
	return &repo{}
}

const (
	collectionName string = "posts"
)

func createClient(ctx context.Context) *firestore.Client {
	sa := option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalf("Failed to create a Firebase App: %v", err)
		return nil
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Failed to create a Firestore client: %v", err)
		return nil
	}

	return client
}

// Save saves a new post into collection
func (rp *repo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()

	_, _, err := client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID":    post.ID,
		"Title": post.Title,
		"Text":  post.Text,
	})

	if err != nil {
		log.Fatalf("Failed adding a new post: %v", err)
		return nil, err
	}

	return post, nil
}

// FindAll returns all posts from the collection
func (rp *repo) FindAll() ([]entity.Post, error) {
	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()

	var posts []entity.Post = make([]entity.Post, 0)

	iter := client.Collection(collectionName).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("failed to iterate the list of posts: %v", err)
			return nil, err
		}

		post := entity.Post{
			ID:    doc.Data()["ID"].(int64),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}

		posts = append(posts, post)
	}

	return posts, nil
}

// Delete deletes a post from database
func (rp *repo) Delete(post *entity.Post) error {
	return nil
}
