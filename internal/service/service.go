package service

import (
	"authentication-service/internal/domain"
	"authentication-service/internal/repository"
)

type Auth interface {
	Login(credentials domain.Credentials) (string, error)
	Auth(token string) (string, error)
}

type Service struct {
	Auth
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Auth: NewAuthService(repos.Auth),
	}
}
