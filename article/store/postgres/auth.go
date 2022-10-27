package postgres

import (
	"context"
	"errors"
	"fmt"
	"github.com/sumit-dhull-97/assignment/article/model"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Article struct {
	DB *pgxpool.Pool
}

func (a *Article) Create(ctx *context.Context, article *model.Article) error {

	_, err := a.DB.Exec(*ctx, "INSERT INTO articles (id, title, user_id, script, hashtags, published, created) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		article.ID, article.Title, article.UserID, article.Script, article.Hashtags, article.Published, article.Created)

	if err != nil {
		fmt.Println(err, article)
		return errors.New("failed to create article")
	}

	return nil
}

func (a *Article) Read(ctx *context.Context, id string) (*model.Article, error) {
	article := model.Article{}

	err := a.DB.QueryRow(*ctx, "SELECT id, title, user_id, script, hashtags, published, created FROM articles WHERE id = $1", id).
		Scan(&article.ID, &article.Title, &article.UserID, &article.Script, &article.Hashtags, &article.Published, &article.Created)

	if err != nil {
		fmt.Println(err, article)
		return nil, errors.New("failed to read article")
	}

	return &article, nil
}

func (a *Article) ReadAll(ctx *context.Context, userId string) ([]model.Article, error) {
	articles := make([]model.Article, 0, 1)

	rows, err := a.DB.Query(*ctx, "SELECT id, title, user_id, script, hashtags, published, created FROM articles WHERE user_id = $1", userId)
	if err != nil {
		fmt.Println(err, userId)
		return nil, errors.New("failed to read articles")
	}

	for rows.Next() {
		article := model.Article{}
		if err := rows.Scan(&article.ID, &article.Title, &article.UserID, &article.Script, &article.Hashtags, &article.Published, &article.Created); err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}

	return articles, nil
}

func (a *Article) Update(ctx *context.Context, article *model.Article) error {

	_, err := a.DB.Exec(*ctx, "UPDATE articles SET title=$2, user_id=$3, script=$4, hashtags=$5, published=$6, created=$7 WHERE id=$1",
		article.ID, article.Title, article.UserID, article.Script, article.Hashtags, article.Published, article.Created)

	if err != nil {
		fmt.Println(err, article)
		return errors.New("failed to update article")
	}

	return nil
}

func (a *Article) Delete(ctx *context.Context, id string) error {

	_, err := a.DB.Exec(*ctx, "DELETE FROM articles WHERE id = $1", id)

	if err != nil {
		fmt.Println(err)
		return errors.New("failed to delete article")
	}

	return nil
}

func GetDBConnection(ctx *context.Context) *pgxpool.Pool {
	host := os.Getenv("DD_DB_HOST")
	if host == "" {
		host = "127.0.0.1"
	}

	const connectionString = "postgres://goland:goland@%s:5432/postgres?sslmode=disable"

	conString := fmt.Sprintf(connectionString, host)

	var dbPool *pgxpool.Pool
	var err error
	for i := 1; i < 8; i++ {
		log.Printf("trying to connect to the db server (attempt %d)...\n", i)
		dbPool, err = pgxpool.Connect(*ctx, fmt.Sprintf(conString, host))
		if err == nil {
			break
		}
		log.Printf("got error: %v\n", err)

		time.Sleep(time.Duration(i*i) * time.Second)
	}

	if dbPool == nil {
		log.Fatalln("could not connect to the database")
	}

	db, err := dbPool.Acquire(*ctx)
	if err != nil {
		log.Fatalf("failed to get connection on startup: %v\n", err)
	}
	if err := db.Conn().Ping(*ctx); err != nil {
		log.Fatalln(err)
	}

	db.Release()

	return dbPool
}
