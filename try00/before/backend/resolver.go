package backend

import (
	"context"
	"fmt"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateItem(ctx context.Context, input ItemInput) (string, error) {
	fmt.Printf("input: %#v\n", input)
	return "item-id-01", nil
}
func (r *mutationResolver) CreateUser(ctx context.Context, input UserInput) (string, error) {
	fmt.Printf("input: %#v\n", input)
	return "user-id-01", nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Item(ctx context.Context, id string) (*Item, error) {
	fmt.Printf("id: %s\n", id)
	return &Item{
		ID:   "item-id-01",
		Name: "item01",
	}, nil
}
func (r *queryResolver) User(ctx context.Context, id string) (*User, error) {
	fmt.Printf("id: %s\n", id)
	return &User{
		ID:   "user-id-01",
		Name: "user01",
	}, nil
}
