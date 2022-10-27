package store

import (
	"context"
	"github.com/sumit-dhull-97/assignment/auth/model"
)

type User interface {
	Create(ctx *context.Context, user *model.User) error
	Read(id string) (*model.User, error)
	Update(user *model.User) error
	Delete(id string) error
}
