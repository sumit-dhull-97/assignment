package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	model2 "github.com/sumit-dhull-97/assignment/article/model"
	"github.com/sumit-dhull-97/assignment/article/store/postgres"

	"github.com/sumit-dhull-97/assignment/article/graph/generated"
	"github.com/sumit-dhull-97/assignment/article/graph/model"
)

// Post is the resolver for the post field.
func (r *mutationResolver) Post(ctx context.Context, input model.ArticleInput) (*model.Article, error) {
	pg := postgres.GetDBConnection(&ctx)
	article := &postgres.Article{DB: pg}

	output := &model2.Article{ID: input.ID, UserID: input.UserID, Script: input.Script, Hashtags: input.Hashtags, Created: "2022-10-10"}

	article.Create(&ctx, output)
	//article.Delete(&ctx, input.ID)
	article.ReadAll(&ctx, "scooby")
	return nil, nil
}

// Delete is the resolver for the delete field.
func (r *mutationResolver) Delete(ctx context.Context, input model.DeleteInput) (string, error) {
	panic(fmt.Errorf("not implemented: Delete - delete"))
}

// GetAll is the resolver for the getAll field.
func (r *queryResolver) GetAll(ctx context.Context, input model.GetAllInput) ([]*model.Article, error) {
	panic(fmt.Errorf("not implemented: GetAll - getAll"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
