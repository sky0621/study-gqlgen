package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/gqlerror"

	"github.com/99designs/gqlgen/handler"
	study_gqlgen "github.com/sky0621/study-gqlgen"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(
		study_gqlgen.NewExecutableSchema(study_gqlgen.Config{Resolvers: &study_gqlgen.Resolver{}}),
		handler.ErrorPresenter(
			func(ctx context.Context, e error) *gqlerror.Error {
				//if myError, ok := e.(MyError); ok {
				//	return gqlerror.ErrorPathf(graphql.GetResolverContext(ctx).Path(), "Eeek!")
				//}

				return graphql.DefaultErrorPresenter(ctx, e)
			},
		),
		handler.RequestMiddleware(
			func(ctx context.Context, next func(ctx context.Context) []byte) []byte {
				return next(ctx)
			},
		),
		handler.ResolverMiddleware(
			func(ctx context.Context, next graphql.Resolver) (res interface{}, err error) {
				return next(ctx)
			},
		),
	))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
