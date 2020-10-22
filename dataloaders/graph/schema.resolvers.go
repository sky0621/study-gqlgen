package graph

import (
	"context"
	"log"

	"github.com/sky0621/study-gqlgen/dataloaders/graph/generated"
	"github.com/sky0621/study-gqlgen/dataloaders/graph/model"
)

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

func (r *userResolver) Todos(ctx context.Context, obj *model.User) ([]*model.Todo, error) {
	if obj == nil {
		return []*model.Todo{}, nil
	}
	return For(ctx).TodosByUserIDs.Load(obj.ID)
}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	if obj == nil {
		return nil, nil
	}
	return For(ctx).UsersByIDs.Load(obj.UserID)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
