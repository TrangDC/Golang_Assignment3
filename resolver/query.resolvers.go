package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"intern_traning/ent"
	graphql1 "intern_traning/graphql"
)

// GetNews is the resolver for the GetNews field.
func (r *queryResolver) GetNews(ctx context.Context, id string) (*ent.NewsResponse, error) {
	return r.serviceRegistry.News().GetNews(ctx, id)
}

// GetAllNews is the resolver for the GetAllNews field.
func (r *queryResolver) GetAllNews(ctx context.Context, pagination *ent.PaginationInput, filter *ent.NewsFilter, freeWord *ent.NewsFreeWord, orderBy *ent.NewsOrder) (*ent.NewsResponseGetAll, error) {
	return r.serviceRegistry.News().GetAllNews(ctx, pagination, filter, orderBy, freeWord)
}

// GetPreRequest is the resolver for the GetPreRequest field.
func (r *queryResolver) GetPreRequest(ctx context.Context) (string, error) {
	return "", nil
}

// Query returns graphql1.QueryResolver implementation.
func (r *Resolver) Query() graphql1.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
