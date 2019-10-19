package utils

import (
	"./token"
)

var (
	t *Token
)

func init() {
	// FIXME: configurable
	t = token.UseJwtToken("MySecretKey")
}

func GenerateToken(userId, userName string) string {
	return t.GenerateToken(userId, userName)
}

func ExplainToken(token string, userId, userName *string) error {
	return t.ExplainToken(token, userId, userName)
}
