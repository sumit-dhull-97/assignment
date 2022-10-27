package service

import (
	"context"
	"github.com/sumit-dhull-97/assignment/auth/model"
)

type Auth interface {
	Login(ctx *context.Context, input *model.User) (*model.User, error)
	Logout(ctx *context.Context, input *model.User) (string, error)
	CheckSession(ctx *context.Context, input *model.User) (string, error)
	Signup(ctx *context.Context, input *model.User) (*model.User, error)
}
