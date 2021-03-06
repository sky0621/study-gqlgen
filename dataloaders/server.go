package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sky0621/study-gqlgen/dataloaders/graph"
	"github.com/sky0621/study-gqlgen/dataloaders/graph/generated"
)

func main() {
	db := sqlx.MustOpen("sqlite3", "./data.db")
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", graph.Middleware(
		db,
		handler.NewDefaultServer(
			generated.NewExecutableSchema(
				generated.Config{
					Resolvers: &graph.Resolver{
						DB: db,
					},
				},
			),
		)),
	)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
