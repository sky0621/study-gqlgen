package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/sky0621/study-gqlgen/types/graph/model"
)

func (r *queryResolver) NonNullTypes(ctx context.Context) ([]*model.NonNullType, error) {
	var results []*model.NonNullType
	results = append(results, &model.NonNullType{
		MapsNonNull:     []map[string]interface{}{nil},
		AnyTypesNonNull: []interface{}{nil},
		ObjectsNonNull:  []*model.Object{nil},
	})
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) NullableTypes(ctx context.Context) ([]*model.NullableType, error) {
	var results []*model.NullableType
	results = append(results, &model.NullableType{})
	panic(fmt.Errorf("not implemented"))
}
