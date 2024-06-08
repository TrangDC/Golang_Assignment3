package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"intern_traning/ent"
	graphql1 "intern_traning/graphql"
)

// CreateNews is the resolver for the CreateNews field.
func (r *mutationResolver) CreateNews(ctx context.Context, input ent.NewNewsInput) (*ent.NewsResponse, error) {
	return r.serviceRegistry.News().CreateNews(ctx, input)
}

// UpdateNews is the resolver for the UpdateNews field.
func (r *mutationResolver) UpdateNews(ctx context.Context, id string, input ent.UpdateNewsInput) (*ent.NewsResponse, error) {
	return r.serviceRegistry.News().UpdateNews(ctx, id, input)
}

// DeleteNews is the resolver for the DeleteNews field.
func (r *mutationResolver) DeleteNews(ctx context.Context, id string) (bool, error) {
	err := r.serviceRegistry.News().DeleteNews(ctx, id)
	if err != nil {
		return false, err
	}
	return true, nil
}

// CreatePreRequest is the resolver for the CreatePreRequest field.
func (r *mutationResolver) CreatePreRequest(ctx context.Context, input string) (string, error) {
	return "", nil
}

// Mutation returns graphql1.MutationResolver implementation.
func (r *Resolver) Mutation() graphql1.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
