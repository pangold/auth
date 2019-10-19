package cache

import (
	"fmt"
)

type Cache interface {
	HasCacheKey(service, action, key string) bool
	GetCacheValue(service, action, key string, value *string) error
	SetCacheValue(service, action, key, value string, timeout int) error
}

func UseSimpleCache() *Cache {
	fmt.Println("Using Simple Cache")
	return &SimpleCache{}
}

func UseRedisCache() *Cache {
	fmt.Println("Using Redis Cache")
	return &RedisCache{}
}

