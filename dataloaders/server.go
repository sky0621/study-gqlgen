package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sky0621/study-gqlgen/dataloaders/graph"
	"github.com/sky0621/study-gqlgen/dataloaders/graph/generated"
)

func main() {
	db := sqlx.MustOpen("sqlite3", "./data.db")

	// 初回起動時のみ
	//setup(db)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		DB: db,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func setup(db *sqlx.DB) {
	db.MustExec(schema)

	tx := db.MustBegin()

	users := []string{"taro", "jiro", "saburo", "siro", "goro", "rokuro"}
	for idx, user := range users {
		tx.MustExec(insertUser, idx+1, user)
	}

	todoID := 1
	for idx, _ := range users {
		for j := 0; j < 3; j++ {
			tx.MustExec(insertTodo, todoID, fmt.Sprintf("ToDo%2d", todoID), idx+1)
			todoID++
		}
	}

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
}

var schema = `
CREATE TABLE todo (
    id integer,
    task text,
    user_id integer
);

CREATE TABLE user (
    id integer,
    name text
);
`

var insertTodo = `INSERT INTO todo VALUES($1, $2, $3)`

var insertUser = `INSERT INTO user VALUES($1, $2)`
