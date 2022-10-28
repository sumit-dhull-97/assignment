package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgryski/trifles/uuid"
	"github.com/sumit-dhull-97/assignment/article/model"
	"github.com/sumit-dhull-97/assignment/article/store"
	"github.com/tidwall/gjson"
	"golang.org/x/net/context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type ArticleService struct {
	Store store.Article
}

func (a *ArticleService) Post(ctx *context.Context, input *model.Article, sessionCred string) (*model.Article, error) {
	err := checkCredentials(input.UserID, sessionCred)
	if err != nil {
		return nil, err
	}

	input.Published = time.Now().String()
	if input.ID == "" {
		input.ID = uuid.UUIDv4()
		input.Created = input.Published

		err = a.Store.Create(ctx, input)

	} else {
		err = a.Store.Update(ctx, input)
	}

	if err != nil {
		return nil, err
	}

	return input, nil
}

func (a *ArticleService) GetAll(ctx *context.Context, userId string, sessionCred string) ([]model.Article, error) {
	err := checkCredentials(userId, sessionCred)
	if err != nil {
		return nil, err
	}

	return a.Store.ReadAll(ctx, userId)
}

func (a *ArticleService) Delete(ctx *context.Context, input *model.Article, sessionCred string) (string, error) {
	err := checkCredentials(input.UserID, sessionCred)
	if err != nil {
		return "", err
	}

	err = a.Store.Delete(ctx, input.ID)
	if err != nil {
		return "", err
	}

	return "DELETED", nil
}

func checkCredentials(userId, sessionCred string) error {
	var query = `
		{
		  checkSession (
			  input: {
				 userId: "%s"
				 sessionCred: "%s" 
			  }
		  ) 
		}
        `
	query = fmt.Sprintf(query, userId, sessionCred)

	jsonData := map[string]string{
		"query": query,
	}
	jsonValue, _ := json.Marshal(jsonData)

	authHost := os.Getenv("AUTH_HOST")
	if authHost == "" {
		authHost = "localhost"
	}

	authPort := os.Getenv("AUTH_PORT")
	if authPort == "" {
		authPort = "8002"
	}

	url := fmt.Sprintf("http://%s:%s/query", authHost, authPort)

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{Timeout: time.Second * 10}
	response, err := client.Do(request)
	defer response.Body.Close()
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	data, _ := io.ReadAll(response.Body)
	var dataString = string(data)

	val := gjson.Get(dataString, "data.checkSession").String()

	if val != "OPEN" {
		log.Println(query)
		return errors.New("wrong Credentials")
	}

	return nil
}
