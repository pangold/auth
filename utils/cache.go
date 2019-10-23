package utils

import (
	"./cache"
)

var (
	c cache.Cache
)

func init() {
	c = cache.UseSimpleCache()
}

func HasCacheKey(service, key string) bool {
	return c.HasCacheKey(service, key)
}

func GetCacheValue(service, key string, vtype interface{}) (interface{}, error) {
	return c.GetCacheValue(service, key, vtype)
}

func SetCacheValue(service, key string, value interface{}, expire int) error {
	return c.SetCacheValue(service, key, value, expire)
}
