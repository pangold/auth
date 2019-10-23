package utils

import (
	"./token"
)

var (
	t token.Token
)

func init() {
	// FIXME: configurable
	t = token.UseJwtToken("MySecretKey")
}

func GenerateToken(userId, userName string, expire int) (string, error) {
	return t.GenerateToken(userId, userName, expire)
}

func ExplainToken(token string, userId, userName *string) error {
	return t.ExplainToken(token, userId, userName)
}

func ResetToken(token string) error {
	return t.ResetToken(token)
}
