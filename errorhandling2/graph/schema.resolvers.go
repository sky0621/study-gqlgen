package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/sky0621/study-gqlgen/errorhandling2/graph/generated"
	"github.com/sky0621/study-gqlgen/errorhandling2/graph/model"
)

type ErrorCode string

const (
	ErrorCodeRequired            ErrorCode = "1001"
	ErrorCodeUnexpectedSituation ErrorCode = "9999"
)

type AppError struct {
	Code ErrorCode
	Msg  string
}

func (e AppError) Error() string {
	return fmt.Sprintf("[%s]%s", e.Code, e.Msg)
}

func (r *queryResolver) ErrorPresenter(ctx context.Context) ([]*model.Todo, error) {
	return nil, AppError{
		Code: ErrorCodeRequired,
		Msg:  "text is none",
	}
}

func (r *queryResolver) PanicHandler(ctx context.Context) ([]*model.Todo, error) {
	panic("unexpected situation")
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
