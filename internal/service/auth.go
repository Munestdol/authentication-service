package service

import (
	config "authentication-service/configs"
	"authentication-service/internal/domain"
	"authentication-service/internal/repository"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/rs/zerolog/log"
)

type AuthService struct {
	repo  repository.Auth
	token string
}

func NewAuthService(repo repository.Auth, cfg *config.Config) *AuthService {
	return &AuthService{repo: repo, token: cfg.Token}
}

func (s *AuthService) Login(creds domain.Credentials) error {
	return s.repo.Login(creds)
}

func (s *AuthService) GetToken(creds domain.Credentials) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &domain.Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.token))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (s *AuthService) Auth(token string) (string, error) {
	tokenString := strings.TrimPrefix(token, "Bearer ")

	claims := &domain.Claims{}
	parsedToken, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) { return []byte(s.token), nil })
	if err != nil || !parsedToken.Valid {
		log.Error().Err(err).Msg("error parse token with claims")
		return "", err
	}

	return claims.Username, nil
}
