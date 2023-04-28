package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.29

import (
	"context"
	"fmt"

	"github.com/phantranhieunhan/s3-assignment/module/friendship/port/graphql/convert"
	"github.com/phantranhieunhan/s3-assignment/module/friendship/port/graphql/model"
)

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented: Users - users"))
}

// Friendships is the resolver for the friendships field.
func (r *queryResolver) Friendships(ctx context.Context) ([]*model.Friendship, error) {
	panic(fmt.Errorf("not implemented: Friendships - friendships"))
}

// Subscriptions is the resolver for the subscriptions field.
func (r *queryResolver) Subscriptions(ctx context.Context) ([]*model.SubscriptionG, error) {
	list, err := r.app.Queries.ListSubscriptions.Handle(ctx)
	if err != nil {
		return []*model.SubscriptionG{}, fmt.Errorf("ListSubscriptions.Handle failed: %v", err)
	}
	return convert.ToSubscriptionG(list), nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }