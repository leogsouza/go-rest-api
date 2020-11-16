package repository

import (
	"database/sql"
	"log"
	"os"

	"github.com/leogsouza/go-rest-api/entity"

	// import to use sqlite library
	_ "github.com/mattn/go-sqlite3"
)

type sqliteRepo struct{}

// NewSQLiteRepository implement post repository using sqlite engine
func NewSQLiteRepository() PostRepository {
	os.Remove("./posts.db")

	db, err := sql.Open("sqlite3", "./posts.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
		create table posts (id integer not null primary key, title text, txt text);
		delete from posts;
	`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
	}

	return &sqliteRepo{}
}

func (sqlRepo *sqliteRepo) Save(post *entity.Post) (*entity.Post, error) {
	db, err := sql.Open("sqlite3", "./posts.db")

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	stmt, err := tx.Prepare("insert into posts(id, title, txt) values(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(post.ID, post.Title, post.Text)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	tx.Commit()

	return post, nil

}

func (sqlRepo *sqliteRepo) FindAll() ([]entity.Post, error) {
	db, err := sql.Open("sqlite3", "./posts.db")

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("select * from posts")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var posts []entity.Post = make([]entity.Post, 0)
	var post entity.Post

	for rows.Next() {
		rows.Scan(&post.ID, &post.Title, &post.Text)
		posts = append(posts, post)
	}

	return posts, nil

}

// Delete deletes a post from database
func (sqlRepo *sqliteRepo) Delete(post *entity.Post) error {
	db, err := sql.Open("sqlite3", "./posts.db")

	if err != nil {
		log.Fatal(err)
		return err
	}
	defer db.Close()
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("delete from posts where id =?")
	_, err = stmt.Exec(post.ID)

	if err != nil {
		log.Fatal(err)
		return err
	}

	tx.Commit()

	return nil

}
