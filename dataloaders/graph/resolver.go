package graph

import "github.com/jmoiron/sqlx"

//go:generate go run github.com/99designs/gqlgen

type Resolver struct {
	DB *sqlx.DB
}
