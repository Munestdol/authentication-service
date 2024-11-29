package service

import (
	"authentication-service/internal/domain"
	"authentication-service/internal/repository"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type AuthService struct {
	repo repository.Auth
}

func NewAuthService(repo repository.Auth) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Login(creds domain.Credentials) (string, error) {
	err := s.repo.Login(creds)
	if err != nil {
		return "", err
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &domain.Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	log.Print(token)

	jwtKey := []byte(setFromEnv())

	log.Print(jwtKey)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		log.Error().Err(err).Msg("error creating token")
		return "", err
	}

	return tokenString, nil
}

func (s *AuthService) Auth() []byte {
	return []byte(setFromEnv())
}

func setFromEnv() (jwtKey string) {
	_ = godotenv.Load()
	return os.Getenv("JWT_KEY")
}
