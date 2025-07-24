package infrustructure

import "golang.org/x/crypto/bcrypt"

func GetHashedPassword(plain_pwd string, cost int) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(plain_pwd), cost)
	if err != nil {
		return "", err
	}
	return string(hashed), err
}

func ValidatePassword(stored_hash string, plain_input string) error {
	return bcrypt.CompareHashAndPassword([]byte(stored_hash), []byte(plain_input))
}
