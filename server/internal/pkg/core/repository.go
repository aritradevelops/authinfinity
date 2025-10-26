package core

import "github.com/aritradevelops/authinfinity/server/internal/pkg/db"

type Repository[S Schema] interface {
	List(*ListOpts) (*PaginatedResponse[S], error)
	Create(S) (string, error)
	Update(Filter, S) (bool, error)
	View(Filter, *S) error
	Delete(Filter) (bool, error)
}

// TODO: build for postgres as well
func NewRepository[S Schema](m Model, database db.Database) Repository[S] {
	return NewPostgresRepository[S](m, database)
}
