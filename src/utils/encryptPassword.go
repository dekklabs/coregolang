package utils

import "golang.org/x/crypto/bcrypt"

//EncryptPassword encripta una contraseña
func EncryptPassword(password string) (string, error) {
	cost := 8

	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}

	return string(hash), err
}
