package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/api"

	"github.com/99designs/gqlgen/codegen/config"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/sky0621/study-gqlgen/querycomplexity/graph"
	"github.com/sky0621/study-gqlgen/querycomplexity/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	cfg, err := config.LoadConfigFromDefaultLocations()
	if err != nil {
		log.Fatalln(err)
	}

	if err := api.Generate(cfg, api.AddPlugin(&MyPlugin{})); err != nil {
		log.Fatalln(err)
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	//srv.Use(extension.FixedComplexityLimit(3))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

type MyPlugin struct {
}

func (p *MyPlugin) Name() string {
	return "MyPlugin"
}
