package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"intern_traning/ent"
	graphql1 "intern_traning/graphql"
)

// ID is the resolver for the id field.
func (r *newsResolver) ID(ctx context.Context, obj *ent.News) (string, error) {
	return obj.ID.String(), nil
}

// Status is the resolver for the status field.
func (r *newsResolver) Status(ctx context.Context, obj *ent.News) (ent.NewsStatus, error) {
	return ent.NewsStatus(obj.Status), nil
}

// AuthorID is the resolver for the author_id field.
func (r *newsResolver) AuthorID(ctx context.Context, obj *ent.News) (string, error) {
	return obj.AuthorID.String(), nil
}

// Author is the resolver for the author field.
func (r *newsResolver) Author(ctx context.Context, obj *ent.News) (*ent.User, error) {
	return obj.Edges.AuthorEdge, nil
}

// News returns graphql1.NewsResolver implementation.
func (r *Resolver) News() graphql1.NewsResolver { return &newsResolver{r} }

type newsResolver struct{ *Resolver }
