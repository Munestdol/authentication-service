package service

import (
	"authentication-service/internal/repository"

	config "authentication-service/configs"
)

type Service struct {
}

func NewService(repos *repository.Repository, cfg *config.Config) *Service {
	return &Service{}
}
