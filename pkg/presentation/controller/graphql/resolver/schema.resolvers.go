package presentation

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"context"
	"fmt"

	presentation1 "github.com/octoposprime/op-be-graphql/pkg/presentation/dto"
	presentation "github.com/octoposprime/op-be-graphql/pkg/presentation/dto/model"
)

// Auth is the resolver for the Auth field.
func (r *authOpsResolver) Auth(ctx context.Context, obj *presentation.AuthOps) (bool, error) {
	panic(fmt.Errorf("not implemented: Auth - Auth"))
}

// Auth is the resolver for the auth field.
func (r *mutationResolver) Auth(ctx context.Context) (*presentation.AuthOps, error) {
	panic(fmt.Errorf("not implemented: Auth - auth"))
}

// Auth is the resolver for the auth field.
func (r *queryResolver) Auth(ctx context.Context) (*presentation.AuthOps, error) {
	panic(fmt.Errorf("not implemented: Auth - auth"))
}

// Status is the resolver for the status field.
func (r *subscriptionResolver) Status(ctx context.Context) (<-chan string, error) {
	panic(fmt.Errorf("not implemented: Status - status"))
}

// AuthOps returns presentation1.AuthOpsResolver implementation.
func (r *Resolver) AuthOps() presentation1.AuthOpsResolver { return &authOpsResolver{r} }

// Mutation returns presentation1.MutationResolver implementation.
func (r *Resolver) Mutation() presentation1.MutationResolver { return &mutationResolver{r} }

// Query returns presentation1.QueryResolver implementation.
func (r *Resolver) Query() presentation1.QueryResolver { return &queryResolver{r} }

// Subscription returns presentation1.SubscriptionResolver implementation.
func (r *Resolver) Subscription() presentation1.SubscriptionResolver { return &subscriptionResolver{r} }

type authOpsResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
