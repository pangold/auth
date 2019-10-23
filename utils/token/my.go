package token

import (
	"errors"
	"strings"

	"../cache"
)

type MyToken struct {
	c cache.Cache
}

func UseMyToken(hostName string) Token {
	cc := cache.UseRedigoCache(hostName)
	// if using simple cache, exipre will be invalid
	// cc := cache.UseSimpleCache() 
	return MyToken{c: cc}
}

func (mt MyToken) GenerateToken(userId, userName string, expire int) (string, error) {
	token := GenerateRandomString(64) // key is token
	value := userId + "," + userName
	if err := mt.c.SetCacheValue("token", token, value, expire); err != nil {
		return "", err
	}
	return token, nil
}

func (mt MyToken) ExplainToken(token string, userId, userName *string) error {
	v, err := mt.c.GetCacheValue("token", token, string(""))
	if err != nil {
		return err
	}
	kv := strings.Split(v.(string), ",")
	if len(kv) != 2 {
		return errors.New("invalid value")
	}
	*userId, *userName = kv[0], kv[1]
	return nil
}

func (mt MyToken) ResetToken(token string) error {
	if err := mt.c.ResetCache("token", token); err != nil {
		return err
	}
	return nil
}
