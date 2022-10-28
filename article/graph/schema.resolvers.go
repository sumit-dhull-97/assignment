package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/sumit-dhull-97/assignment/article/graph/generated"
	"github.com/sumit-dhull-97/assignment/article/graph/model"
	model2 "github.com/sumit-dhull-97/assignment/article/model"
)

// Post is the resolver for the post field.
func (r *mutationResolver) Post(ctx context.Context, input model.ArticleInput) (*model.Article, error) {
	article := model2.Article{ID: input.ID, UserID: input.UserID, Script: input.Script, Title: input.Title, Hashtags: input.Hashtags}

	output, err := r.Service.Post(&ctx, &article, input.SessionCred)
	if err != nil {
		return nil, err
	}

	return &model.Article{ID: output.ID, UserID: output.UserID, Script: output.Script, Title: output.Title, Hashtags: output.Hashtags,
		Published: output.Published, Created: output.Created}, nil
}

// Delete is the resolver for the delete field.
func (r *mutationResolver) Delete(ctx context.Context, input model.DeleteInput) (string, error) {
	article := &model2.Article{ID: input.ArticleID, UserID: input.UserID}

	return r.Service.Delete(&ctx, article, input.SessionCred)
}

// GetAll is the resolver for the getAll field.
func (r *queryResolver) GetAll(ctx context.Context, input model.GetAllInput) ([]*model.Article, error) {
	articles, err := r.Service.GetAll(&ctx, input.UserID, input.SessionCred)
	if err != nil {
		return nil, err
	}

	res := make([]*model.Article, 0, 1)
	for _, v := range articles {
		r := &model.Article{ID: v.ID, UserID: v.UserID, Script: v.Script, Title: v.Title, Hashtags: v.Hashtags,
			Published: v.Published, Created: v.Created}
		res = append(res, r)
	}

	return res, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
