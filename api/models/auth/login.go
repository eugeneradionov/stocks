package auth

import "github.com/eugeneradionov/stocks/api/models"

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResp struct {
	User  models.User  `json:"user"`
	Token models.Token `json:"token"`
}
