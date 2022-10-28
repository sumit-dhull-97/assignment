package postgres

import (
	"context"
	"errors"
	"fmt"
	"github.com/sumit-dhull-97/assignment/auth/model"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

type User struct {
	DB *pgxpool.Pool
}

func (a *User) Create(ctx *context.Context, user *model.User) error {

	_, err := a.DB.Exec(*ctx, "INSERT INTO users (id, first_name, last_name, mobile, password, session_id) VALUES ($1, $2, $3, $4, $5, $6)",
		user.ID, user.FirstName, user.LastName, user.Mobile, user.Password, user.SessionCred)

	if err != nil {
		fmt.Println(err, user)
		return errors.New("failed to create user")
	}

	return nil
}

func (a *User) Read(ctx *context.Context, id string) (*model.User, error) {
	user := model.User{}

	err := a.DB.QueryRow(*ctx, "SELECT id, first_name, last_name, mobile, password, session_id FROM users WHERE id = $1", id).
		Scan(&user.ID, &user.FirstName, &user.LastName, &user.Mobile, &user.Password, &user.SessionCred)

	if err != nil {
		fmt.Println(err, user)
		return nil, errors.New("failed to read user")
	}

	return &user, nil
}

func (a *User) Update(ctx *context.Context, user *model.User) error {

	_, err := a.DB.Exec(*ctx, "UPDATE users SET first_name=$2, last_name=$3, mobile=$4, password=$5, session_id=$6 WHERE id=$1",
		user.ID, user.FirstName, user.LastName, user.Mobile, user.Password, user.SessionCred)

	if err != nil {
		fmt.Println(err, user)
		return errors.New("failed to update user")
	}

	return nil
}

func (a *User) Delete(ctx *context.Context, id string) error {

	_, err := a.DB.Exec(*ctx, "DELETE FROM users WHERE id = $1", id)

	if err != nil {
		fmt.Println(err)
		return errors.New("failed to delete user")
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
