package utils

import (
	"gitlab.com/pangold/auth/utils/token"
)

var (
	t token.Token
)

func getToken() token.Token {
	if t == nil {
		t = token.UseJwtToken(GetConfig().Jwt.SecretKey)
	}
	return t
}

func GenerateToken(userId, userName string, expire int) (string, error) {
	return getToken().GenerateToken(userId, userName, expire)
}

func ExplainToken(token string, userId, userName *string) error {
	return getToken().ExplainToken(token, userId, userName)
}

func ResetToken(token string) error {
	return getToken().ResetToken(token)
}
