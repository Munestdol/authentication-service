package repository

import (
	"authentication-service/internal/domain"
	"errors"

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
	if err == nil {
		log.Error().Err(err).Msg("user already exists")
		return errors.New("user already exists")
	}

	query = `INSERT INTO users (username, password)
	VALUES ($1, $2) RETURNING id`

	row := r.db.QueryRow(query, credentials.Username, credentials.Password)
	if err = row.Scan(&id); err != nil {
		log.Error().Err(err).Msg("error for sign up user")
		return err
	}
	return nil
}
