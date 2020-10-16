package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	"github.com/sky0621/study-gqlgen/errorhandling/graph/generated"
	"github.com/sky0621/study-gqlgen/errorhandling/graph/model"
)

func (r *queryResolver) NormalReturn(ctx context.Context) ([]*model.Todo, error) {
	return []*model.Todo{
		{ID: "001", Text: "something1"},
		{ID: "002", Text: "something2"},
	}, nil
}

func (r *queryResolver) ErrorReturn(ctx context.Context) ([]*model.Todo, error) {
	return nil, errors.New("error occurred")
}

func (r *queryResolver) PanicReturn(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("panic occurred"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}
