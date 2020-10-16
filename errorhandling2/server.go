package main

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/sky0621/study-gqlgen/errorhandling2/graph"
	"github.com/sky0621/study-gqlgen/errorhandling2/graph/generated"
)

func main() {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	srv.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
		err := graphql.DefaultErrorPresenter(ctx, e)

		var appErr graph.AppError
		if errors.As(err, &appErr) {
			return &gqlerror.Error{
				Message: appErr.Msg,
				Extensions: map[string]interface{}{
					"code": appErr.Code,
				},
			}
		}
		return err
	})

	srv.SetRecoverFunc(func(ctx context.Context, err interface{}) error {
		return &gqlerror.Error{
			Extensions: map[string]interface{}{
				"code":  graph.ErrorCodeUnexpectedSituation,
				"cause": err,
			},
		}
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
