package utils

import (
	"strconv"

	"gitlab.com/pangold/auth/utils/cache"
)

var (
	c cache.Cache
)

func getCache() cache.Cache {
	if c == nil {
		c = cache.UseRedigoCache(GetConfig().Redis.Host + ":" + strconv.Itoa(GetConfig().Redis.Port))
		//c = cache.UseSimpleCache()
	}
	return c
}

func HasCacheKey(service, key string) bool {
	return getCache().HasCacheKey(service, key)
}

func GetCacheValue(service, key string, vtype interface{}) (interface{}, error) {
	return getCache().GetCacheValue(service, key, vtype)
}

func SetCacheValue(service, key string, value interface{}, expire int) error {
	return getCache().SetCacheValue(service, key, value, expire)
}

func ResetCache(service, key string) error {
	return getCache().ResetCache(service, key)
}
