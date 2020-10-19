package graph

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/sky0621/study-gqlgen/dataloaders/graph/generated"
	"github.com/sky0621/study-gqlgen/dataloaders/graph/model"
)

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	var todos []*model.Todo
	sql := "SELECT * FROM todo"
	log.Print(sql)
	if err := r.DB.SelectContext(ctx, &todos, sql); err != nil {
		log.Print(err)
		return nil, err
	}
	return todos, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	sql := "SELECT * FROM user"
	log.Print(sql)
	if err := r.DB.SelectContext(ctx, &users, sql); err != nil {
		log.Print(err)
		return nil, err
	}
	return users, nil
}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	if obj == nil {
		return nil, nil
	}
	var users []*model.User
	sql := fmt.Sprintf("SELECT * FROM user WHERE id = %d", obj.UserID)
	log.Print(sql)
	if err := r.DB.SelectContext(ctx, &users, sql); err != nil {
		log.Print(err)
		return nil, err
	}
	if len(users) != 1 {
		log.Print("users length is not 1")
		return nil, errors.New("err")
	}
	return users[0], nil
}

func (r *userResolver) Todos(ctx context.Context, obj *model.User) ([]*model.Todo, error) {
	if obj == nil {
		return []*model.Todo{}, nil
	}
	var todos []*model.Todo
	sql := fmt.Sprintf("SELECT * FROM todo WHERE user_id = %d", obj.ID)
	log.Print(sql)
	if err := r.DB.SelectContext(ctx, &todos, sql); err != nil {
		log.Print(err)
		return nil, err
	}
	return todos, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
