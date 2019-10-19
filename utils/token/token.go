package token

import (
	"fmt"
)

type Token interface {
	GenerateToken(userId, userName string) string
	ExplainToken(token string, userId, userName *string) error
}

func UseMyToken(hostName string) *Token {
	fmt.Println("use my token")
	token := &MyToken{}
	token.Connect(hostName)
	return token
}

func UseJwtToken(secretKey string) *Token {
	fmt.Println("use jwt token")
	return &JwtToken{SecretKey: secretKey}
}
