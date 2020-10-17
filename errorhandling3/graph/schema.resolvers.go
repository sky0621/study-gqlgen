package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/sky0621/study-gqlgen/errorhandling3/graph/generated"
	"github.com/sky0621/study-gqlgen/errorhandling3/graph/model"
)

func (r *queryResolver) CustomErrorReturn(ctx context.Context) ([]*model.Todo, error) {
	// 認証エラーを追加
	NewAuthenticationError().AddGraphQLError(ctx)

	// 認可エラーを追加
	NewAuthorizationError().AddGraphQLError(ctx)

	// バリデーションエラーを追加
	NewValidationError("name", "taro").AddGraphQLError(ctx)

	// その他のエラーを追加
	NewInternalServerError().AddGraphQLError(ctx)

	return nil, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
