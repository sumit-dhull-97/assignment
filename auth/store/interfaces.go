package store

import (
	"context"
	"github.com/sumit-dhull-97/assignment/auth/model"
)

type User interface {
	Create(ctx *context.Context, user *model.User) error
	Read(ctx *context.Context, id string) (*model.User, error)
	Update(ctx *context.Context, user *model.User) error
	Delete(ctx *context.Context, id string) error
}
