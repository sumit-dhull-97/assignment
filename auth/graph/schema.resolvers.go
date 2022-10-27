package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/sumit-dhull-97/assignment/auth/graph/generated"
	"github.com/sumit-dhull-97/assignment/auth/graph/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (*model.Login, error) {
	return &model.Login{
		SessionCred: "some",
	}, nil
}

// Signup is the resolver for the signup field.
func (r *mutationResolver) Signup(ctx context.Context, input model.UserInput) (*model.User, error) {
	return &model.User{
		ID:          "fdggfgg",
		FirstName:   input.FirstName,
		LastName:    input.LastName,
		Mobile:      input.Mobile,
		SessionCred: "some",
	}, nil
}

// Logout is the resolver for the logout field.
func (r *mutationResolver) Logout(ctx context.Context, input *model.LogoutInput) (*model.SessionStatus, error) {
	return nil, gqlerror.Errorf("my error")
}

// CheckSession is the resolver for the checkSession field.
func (r *queryResolver) CheckSession(ctx context.Context, input model.CheckSessionInput) (*model.SessionStatus, error) {
	panic(fmt.Errorf("not implemented: CheckSession - checkSession"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
