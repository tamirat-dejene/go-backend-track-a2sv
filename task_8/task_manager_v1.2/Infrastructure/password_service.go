package infrastructure

import (
	"golang.org/x/crypto/bcrypt"
)

type PasswordService struct{}

func (p *PasswordService) GetHashedPassword(plain_pwd string, cost int) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(plain_pwd), cost)
	if err != nil {
		return "", err
	}
	return string(hashed), err
}

func (p *PasswordService) ValidatePassword(stored_hash string, plain_input string) error {
	return bcrypt.CompareHashAndPassword([]byte(stored_hash), []byte(plain_input))
}

func NewPasswordService() *PasswordService {
	return &PasswordService{}
}
