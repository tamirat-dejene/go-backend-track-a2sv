package infrustructure

import (
	domain "t7/taskmanager/Domain"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(user_name string, expiry int, secret string) (string, error) {
	exp := time.Now().Add(time.Hour * time.Duration(expiry)).Unix()
	claims := domain.JWTPayload{
		UserName: user_name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func IsAuthorized(token_str string, secret_key []byte) (string, error) {
	token, err := jwt.Parse(token_str, func(t *jwt.Token) (any, error) {
		return secret_key, nil
	})

	if err != nil || !token.Valid {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", err
	}

	user_name := claims["user_name"].(string)
	return user_name, nil
}