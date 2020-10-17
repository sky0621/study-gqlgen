package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/sky0621/study-gqlgen/errorhandling/graph/generated"
	"github.com/sky0621/study-gqlgen/errorhandling/graph/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
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

func (r *queryResolver) CustomErrorReturn(ctx context.Context) ([]*model.Todo, error) {
	return nil, gqlerror.Errorf("custom error")
}

func (r *queryResolver) CustomErrorReturn2(ctx context.Context) ([]*model.Todo, error) {
	graphql.AddError(ctx, gqlerror.Errorf("add error"))
	graphql.AddErrorf(ctx, "add error2: %s", time.Now().String())
	return nil, nil
}

func (r *queryResolver) CustomErrorReturn3(ctx context.Context) ([]*model.Todo, error) {
	return nil, &gqlerror.Error{
		Extensions: map[string]interface{}{
			"code":  "A00001",
			"field": "text",
			"value": "トイレ掃除",
		},
	}
}

func (r *queryResolver) CustomErrorReturn4(ctx context.Context) ([]*model.Todo, error) {
	return nil, &gqlerror.Error{
		Extensions: map[string]interface{}{
			"errors": []map[string]interface{}{
				{
					"code":  "A00001",
					"field": "text",
					"value": "トイレ掃除",
				},
				{
					"code":  "A00002",
					"field": "text",
					"value": "トイレ掃除",
				},
			},
		},
	}
}

func (r *queryResolver) PanicReturn(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("panic occurred"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
