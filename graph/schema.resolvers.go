package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.44-dev

import (
	"context"
	"fmt"
)

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*User, error) {
	panic(fmt.Errorf("not implemented: Users - users"))
}

// Pets is the resolver for the pets field.
func (r *userResolver) Pets(ctx context.Context, obj *User) ([]*Pet, error) {
	panic(fmt.Errorf("not implemented: Pets - pets"))
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
