package stringutils

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateRandomString(s int) (string, error) {
	b, err := generateRandomBytes(s)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(b), nil
}

func generateRandomBytes(n int) (b []byte, err error) {
	b = make([]byte, n)

	_, err = rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}
