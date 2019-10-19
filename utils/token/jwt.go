package token

import (

)

type JwtToken struct {
	SecretKey string
}

func (jt *JwtToken) GenerateToken(userId, userName string) string {
	return userId + "." + userName
}

func (jt *JwtToken) ExplainToken(token string, userId, userName *string) error {
	return nil
}
