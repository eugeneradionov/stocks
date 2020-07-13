package token

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/eugeneradionov/stocks/api/models"
	string_utils "github.com/eugeneradionov/stocks/api/string-utils"
)

func (srv service) Generate(claims *Claims) (string, error) {
	claims.StandardClaims.ExpiresAt = time.Now().Unix() + int64(srv.expirationTimeSec)

	tokenWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenWithClaims.SignedString([]byte(srv.secretKey))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (srv service) GenerateRefresh() (string, error) {
	return string_utils.GenerateRandomString(models.RefreshTokenLen)
}

func (srv service) parseClaims(token *jwt.Token) *Claims {
	if claims, ok := token.Claims.(*Claims); ok {
		return claims
	}

	return nil
}
