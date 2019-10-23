package cache

import (
	"errors"
)

var (
	storage map[string]interface{}
)

type SimpleCache struct {}


func UseSimpleCache() Cache {
	if storage == nil {
		storage = make(map[string]interface{})
	}
	return SimpleCache{}
}

func (sc SimpleCache) HasCacheKey(service, key string) bool {
	k := generateCacheKey(service, key)
	_, ok := storage[k]
	return ok
}

func (sc SimpleCache) SetCacheValue(service, key string, value interface{}, expire int) error {
	k := generateCacheKey(service, key)
	storage[k] = value
	return nil
}

func (sc SimpleCache) GetCacheValue(service, key string, vtype interface{}) (interface{}, error) {
	k := generateCacheKey(service, key)
	v, ok := storage[k]
	if !ok {
		return nil, errors.New("key " + k + " is not exist")
	}
	return v, nil
}

func (sc SimpleCache) ResetCache(service, key string) error {
	k := generateCacheKey(service, key)
	delete(storage, k)
	return nil
}
