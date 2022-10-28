package service

import (
	"github.com/sumit-dhull-97/assignment/article/model"
	"golang.org/x/net/context"
)

type Article interface {
	Post(ctx *context.Context, input *model.Article, sessionCred string) (*model.Article, error)
	GetAll(ctx *context.Context, userId string, sessionCred string) ([]model.Article, error)
	Delete(ctx *context.Context, input *model.Article, sessionCred string) (string, error)
}
