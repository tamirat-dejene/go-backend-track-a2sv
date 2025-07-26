package domain

import "github.com/golang-jwt/jwt"

type JWTPayload struct {
	UserName string `json:"user_name"`
	jwt.StandardClaims
}
