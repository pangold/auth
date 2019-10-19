package utils

import (
	"time"
	"./cache"
)

var (
	c *Cache
)

func init() {
	c = cache.UseSimpleCache()
}

func HasCacheKey(service, action, key string) bool {
	return c.HasCacheKey(service, action, key)
}

func GetCacheValue(service, action, key string, value *string) error {
	return c.GetCacheValue(service, action, key, value)
}

func SetCacheValue(service, action, key, value string, timeout int) error {
	return c.SetCacheValue(service, action, key, value, timeout)
}
