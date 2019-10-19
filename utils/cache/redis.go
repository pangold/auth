package cache

import (
	"fmt"
)


type RedisCache struct {
	storage map[string]string
}

func (rc RedisCache) HasCacheKey(service, action, key string) bool {
	k := rc.generateCacheKey(service, action, key)
	_, ok := rc.storage[k]
	return ok
}

func (rc RedisCache) GetCacheValue(service, action, key string, value *string) error {
	k := rc.generateCacheKey(service, action, key)
	*value = rc.storage[k]
	return nil
}

func (rc *RedisCache) SetCacheValue(service, action, key, value string, timeout int) error {
	k := rc.generateCacheKey(service, action, key)
	rc.storage[k] = value
	return nil
}

func (RedisCache) generateCacheKey(service, action, key string) string {
	return fmt.Sprintf("%s.%s.%s", service, action, key)
}
