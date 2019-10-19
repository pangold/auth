package cache

import (
	"fmt"
)


type SimpleCache struct {
	storage map[string]string
}

func (sc SimpleCache) HasCacheKey(service, action, key string) bool {
	k := sc.generateCacheKey(service, action, key)
	_, ok := sc.storage[k]
	return ok
}

func (sc RedisCache) GetCacheValue(service, action, key string, value *string) error {
	k := sc.generateCacheKey(service, action, key)
	*value = sc.storage[k]
	return nil
}

func (sc *SimpleCache) SetCacheValue(service, action, key, value string, timeout int) error {
	k := sc.generateCacheKey(service, action, key)
	sc.storage[k] = value
	return nil
}

func (SimpleCache) generateCacheKey(service, action, key string) string {
	return fmt.Sprintf("%s.%s.%s", service, action, key)
}
