package models

type Exchange struct {
	Code     string `json:"code"`
	Currency string `json:"currency"`
	Name     string `json:"name"`
}
