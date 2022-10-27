package service

import "github.com/sumit-dhull-97/assignment/auth/graph/model"

type Auth interface {
	Login(input model.LoginInput) (model.Login, error)
	Logout(input model.LogoutInput) (model.SessionStatus, error)
	CheckSession(input model.CheckSessionInput) (model.SessionStatus, error)
	Signup(input model.UserInput) (model.User, error)
}
