package system

import (
	"errors"
	"gitlab.com/pangold/auth/middleware"
	"gitlab.com/pangold/auth/utils"
	"strings"
)

type CacheToken struct {
	cache middleware.Cache
}

func NewCacheToken(cache middleware.Cache) *CacheToken {
	return &CacheToken{
		cache: cache,
	}
}

func (this CacheToken) GenerateToken(userId, userName string, expire int) (string, error) {
	token := utils.GenerateRandomString(64)
	value := userId + "," + userName
	if err := this.cache.SetCacheValue("token", token, value, expire); err != nil {
		return "", err
	}
	return token, nil
}

func (this CacheToken) ExplainToken(token string, userId, userName *string) error {
	v, err := this.cache.GetCacheValue("token", token, string(""))
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

func (this CacheToken) ResetToken(token string) error {
	if err := this.cache.ResetCacheKey("token", token); err != nil {
		return err
	}
	return nil
}
