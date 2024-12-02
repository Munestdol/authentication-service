package service

import (
	config "authentication-service/configs"
	"authentication-service/internal/domain"
	"authentication-service/internal/repository"
)

type Auth interface {
	Login(creds domain.Credentials) error
	GetToken(creds domain.Credentials) (string, error)
	Auth(token string) (string, error)
}

type Service struct {
	Auth
}

func NewService(repos *repository.Repository, cfg *config.Config) *Service {
	return &Service{
		Auth: NewAuthService(repos.Auth, cfg),
	}
}
