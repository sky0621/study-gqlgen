package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql"
	"github.com/sky0621/study-gqlgen/sample/graph/model"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/sky0621/study-gqlgen/sample/graph"
	"github.com/sky0621/study-gqlgen/sample/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	resolver := &graph.Resolver{}

	cfg := generated.Config{Resolvers: resolver}
	cfg.Directives.HasRole = func(ctx context.Context, obj interface{}, next graphql.Resolver, role model.Role) (res interface{}, err error) {
		//exts := graphql.GetExtensions(ctx)
		//fmt.Println(exts)
		//fldCtx := graphql.GetFieldContext(ctx)
		//fmt.Println(fldCtx)
		opeCtx := graphql.GetOperationContext(ctx)
		fmt.Println("OperationName: " + opeCtx.OperationName)
		fmt.Println("======================================")
		fmt.Println(opeCtx.RawQuery)
		fmt.Println("======================================")
		fmt.Println("[Operation]")
		//fmt.Println(opeCtx.Operation)
		fmt.Println(opeCtx.Operation.Name)
		fmt.Println(opeCtx.Operation.SelectionSet)
		fmt.Println(opeCtx.Operation.Directives)
		fmt.Println("[Stats]")
		fmt.Println(opeCtx.Stats)
		fmt.Println("[Variables]")
		fmt.Println(opeCtx.Variables)
		//fmt.Println(opeCtx)
		//pathCtx := graphql.GetPathContext(ctx)
		//fmt.Println(pathCtx)
		return next(ctx)
	}
	cfg.Directives.Abc = func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
		fmt.Println(obj)
		return next(ctx)
	}

	schema := generated.NewExecutableSchema(cfg)
	srv := handler.NewDefaultServer(schema)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
