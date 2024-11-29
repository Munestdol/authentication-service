package repository

import (
	"authentication-service/internal/domain"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) Login(credentials domain.Credentials) error {
	var id int
	query := `SELECT id FROM users WHERE username = $1 AND password = $2`
	err := r.db.Get(&id, query, &credentials.Username, &credentials.Password)
	if err != nil {
		log.Error().Err(err).Msg("Wrong username or password")
		return err
	}
	return nil
}
