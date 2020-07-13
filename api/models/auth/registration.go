package auth

import "github.com/eugeneradionov/stocks/api/models"

type RegistrationReq struct {
	Name     string `json:"name"     validate:"required,gte=2,lte=70"`
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required,gte=8,lte=32"`
}

type RegistrationResp struct {
	User  models.User  `json:"user"`
	Token models.Token `json:"token"`
}
