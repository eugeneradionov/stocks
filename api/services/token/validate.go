package token

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/eugeneradionov/stocks/api/models"
	"github.com/google/uuid"
)

var (
	ErrTokenInvalid       = errors.New("token is invalid")
	ErrTokenInBlackList   = errors.New("token in black list")
	ErrTokenNotFound      = errors.New("token not found")
	ErrTokenClaimsInvalid = errors.New("token claims is invalid")
)

func (srv service) Validate(accessToken string) (*models.LoginSession, error) {
	token, err := srv.parseJWT(accessToken)
	if err != nil {
		return nil, err
	}

	claims := srv.parseClaims(token)
	if claims == nil || !token.Valid {
		return nil, ErrTokenInvalid
	}

	if err := srv.checkBlackList(claims.TokenID); err != nil {
		return nil, err
	}

	return &models.LoginSession{
		UserID:  claims.UserID,
		TokenID: claims.TokenID,
	}, nil
}

func (srv service) checkBlackList(tokenID uuid.UUID) error {
	res, err := srv.rds.Get(srv.TokenBlackListKey(tokenID))
	if err != nil {
		return err
	}

	if res != nil {
		return ErrTokenInBlackList
	}

	return nil
}

func (srv service) parseJWT(tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(srv.secretKey), nil
	})

	if token == nil {
		return nil, ErrTokenInvalid
	}

	switch validationErr := err.(type) {
	case *jwt.ValidationError:
		if validationErr.Errors|(jwt.ValidationErrorExpired) != jwt.ValidationErrorExpired {
			return nil, ErrTokenInvalid
		}
	case nil:
	default:
		return nil, err
	}

	return token, nil
}
