package password

import (
	"golang.org/x/crypto/bcrypt"
)

func (srv service) EncryptPassword(password, salt string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(srv.saltedPassword(password, salt)), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (srv service) CompareHashAndPassword(password, salt, hash string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(srv.saltedPassword(password, salt))); err != nil {
		return false
	}

	return true
}
