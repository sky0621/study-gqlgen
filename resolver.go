package study_gqlgen

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/vektah/gqlparser/gqlerror"

	"github.com/sky0621/study-gqlgen/model"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	todos []*model.Todo
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Todo() TodoResolver {
	return &todoResolver{r}
}

type mutationResolver struct{ *Resolver }

// TODOを作成する。
func (r *mutationResolver) CreateTodo(ctx context.Context, input NewTodo) (*model.Todo, error) {
	if len(input.UserID) != 3 {
		return nil, &gqlerror.Error{
			Message: "length is only 3",
			Locations: []gqlerror.Location{
				{
					Column: 1,
					Line:   0,
				},
			},
			Extensions: map[string]interface{}{
				"error_code": "400-000-0001-00105-00000",
			},
		}
	}
	todo := &model.Todo{
		ID:     fmt.Sprintf("T%d", rand.Int()),
		Text:   input.Text,
		Done:   false,
		UserID: input.UserID,
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

type todoResolver struct{ *Resolver }

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*User, error) {
	return &User{
		ID:   obj.UserID,
		Name: "user " + obj.UserID,
	}, nil
}
