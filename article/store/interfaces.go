package store

import (
	"context"
	"github.com/sumit-dhull-97/assignment/article/model"
)

type Article interface {
	Create(ctx *context.Context, user *model.Article) error
	Read(ctx *context.Context, id string) (*model.Article, error)
	ReadAll(ctx *context.Context, userId string) ([]model.Article, error)
	Update(ctx *context.Context, user *model.Article) error
	Delete(ctx *context.Context, id string) error
}
