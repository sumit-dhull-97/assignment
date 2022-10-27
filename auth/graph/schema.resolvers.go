package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/sumit-dhull-97/assignment/auth/graph/generated"
	"github.com/sumit-dhull-97/assignment/auth/graph/model"
	model2 "github.com/sumit-dhull-97/assignment/auth/model"
)

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (*model.Login, error) {
	user := model2.User{ID: input.UserID, Password: input.Password}
	output, err := r.Service.Login(&ctx, &user)

	if err != nil {
		return nil, err
	}

	return &model.Login{SessionCred: output.SessionCred}, err
}

// Signup is the resolver for the signup field.
func (r *mutationResolver) Signup(ctx context.Context, input model.UserInput) (*model.User, error) {
	user := model2.User{FirstName: input.FirstName, LastName: input.LastName, Mobile: input.Mobile, Password: input.Password}

	output, err := r.Service.Signup(&ctx, &user)
	if err != nil {
		return nil, err
	}

	return &model.User{ID: output.ID, FirstName: output.FirstName, LastName: output.LastName, Mobile: output.Mobile, SessionCred: output.SessionCred}, nil
}

// Logout is the resolver for the logout field.
func (r *mutationResolver) Logout(ctx context.Context, input *model.LogoutInput) (*model.SessionStatus, error) {
	user := model2.User{ID: input.UserID, SessionCred: input.SessionCred}

	output, err := r.Service.Logout(&ctx, &user)
	if err != nil {
		return nil, err
	}

	var ss = model.SessionStatus(output)
	return &ss, nil
}

// CheckSession is the resolver for the checkSession field.
func (r *queryResolver) CheckSession(ctx context.Context, input model.CheckSessionInput) (*model.SessionStatus, error) {
	user := model2.User{ID: input.UserID, SessionCred: input.SessionCred}

	output, err := r.Service.CheckSession(&ctx, &user)
	if err != nil {
		return nil, err
	}

	var ss = model.SessionStatus(output)
	return &ss, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
