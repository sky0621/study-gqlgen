package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/sky0621/study-gqlgen/querycomplexity/graph/generated"
	"github.com/sky0621/study-gqlgen/querycomplexity/graph/model"
)

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return []*model.Todo{
		{
			ID:   "001",
			Text: "something1",
			User: &model.User{
				ID:   "a01",
				Name: "taro",
				Todos: []*model.Todo{
					{
						ID:   "001",
						Text: "something1",
						User: &model.User{
							ID:   "a01",
							Name: "taro",
							Todos: []*model.Todo{
								{
									ID:   "001",
									Text: "something1",
									User: &model.User{
										ID:    "a01",
										Name:  "taro",
										Todos: []*model.Todo{},
									},
								},
							},
						},
					},
				},
			},
		},
	}, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	// FIXME:
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
