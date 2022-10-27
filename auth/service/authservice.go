package service

import (
	"errors"
	"github.com/dgryski/trifles/uuid"
	"github.com/sumit-dhull-97/assignment/auth/model"
	"github.com/sumit-dhull-97/assignment/auth/store"
	"golang.org/x/net/context"
	"log"
)

type AuthService struct {
	Store store.User
}

func (a *AuthService) Login(ctx *context.Context, input *model.User) (*model.User, error) {
	user, err := a.Store.Read(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	if input.Password != user.Password {
		return nil, errors.New("wrong password")
	}

	user.SessionCred = uuid.UUIDv4()
	err = a.Store.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (a *AuthService) Logout(ctx *context.Context, input *model.User) (string, error) {
	user, err := a.Store.Read(ctx, input.ID)
	if err != nil {
		return "", err
	}

	if user.SessionCred != input.SessionCred {
		log.Println("logout failed " + user.SessionCred + " - " + input.SessionCred)
		return "", errors.New("wrong session token")
	}

	user.SessionCred = uuid.UUIDv4()
	err = a.Store.Update(ctx, user)
	if err != nil {
		return "", err
	}

	return model.TERMINATED_SESSION, nil
}

func (a *AuthService) CheckSession(ctx *context.Context, input *model.User) (string, error) {
	user, err := a.Store.Read(ctx, input.ID)
	if err != nil {
		return "", err
	}

	if user.SessionCred != input.SessionCred {
		log.Println("session check failed " + user.SessionCred + " - " + input.SessionCred)
		return model.TERMINATED_SESSION, nil
	}

	return model.OPEN_SESSION, nil
}

func (a *AuthService) Signup(ctx *context.Context, input *model.User) (*model.User, error) {
	input.ID = uuid.UUIDv4()
	input.SessionCred = uuid.UUIDv4()

	err := a.Store.Create(ctx, input)
	if err != nil {
		log.Println("sign up failed")
		return nil, err
	}

	return input, nil
}

//
//func (a *AuthService) Logout(input *model.LogoutInput) (*model.SessionStatus, error) {
//
//}
//
//func (a *AuthService) CheckSession(input *model.CheckSessionInput) (*model.SessionStatus, error) {
//
//}
//
//func (a *AuthService) Signup(input *model.UserInput) (*model.User, error) {
//
//}
