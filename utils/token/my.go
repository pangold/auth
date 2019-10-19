package token

import (
	"../redis"
)

type MyToken struct {
	redis.MyRedis
}

func (mt *MyToken) GenerateToken(userId, userName string) string {
	return userId + "." + userName
}

func (mt *MyToken) ExplainToken(token string, userId, userName *string) error {
	return nil
}
