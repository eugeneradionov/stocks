package password

import string_utils "github.com/eugeneradionov/stocks/api/string-utils"

const DefaultSaltSize = 128

func (srv service) GenerateSalt() (string, error) {
	return string_utils.GenerateRandomString(DefaultSaltSize)
}

func (srv service) saltedPassword(password, salt string) string {
	return password + salt
}
