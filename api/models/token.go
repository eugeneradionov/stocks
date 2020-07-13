package models

const RefreshTokenLen = 64

type Token struct {
	Access  string `json:"access_token"`
	Refresh string `json:"refresh_token"`
}
