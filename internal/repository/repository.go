package repository

import (
	"authentication-service/internal/domain"

	"github.com/jmoiron/sqlx"
)

type Auth interface {
	Login(credentials domain.Credentials) error
}

type Repository struct {
	Auth
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Auth: NewAuthPostgres(db),
	}
}
