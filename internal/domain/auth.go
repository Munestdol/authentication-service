package domain

import "github.com/dgrijalva/jwt-go"

type Credentials struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
