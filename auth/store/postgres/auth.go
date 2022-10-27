package postgres

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/sumit-dhull-97/assignment/auth/graph/model"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Auth struct {
	DB *pgxpool.Pool
}

func (a *Auth) Login(ctx *context.Context, input *model.LoginInput) (*model.Login, error) {
	var pass string
	err := a.DB.QueryRow(*ctx, "SELECT password FROM users WHERE id = $1", input.UserID).Scan(&pass)
	if err != nil {
		log.Println(err, input)
		return nil, err
	}

	session, _ := uuid.NewV1()

	_, err = a.DB.Exec(*ctx, "UPDATE users SET session_id = $1 WHERE id = $2", session.String(), input.UserID)
	if err != nil {
		log.Println(err, session, input)
		return nil, err
	}

	return &model.Login{
		SessionCred: session.String(),
	}, nil
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
