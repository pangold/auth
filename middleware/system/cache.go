package system

import (
	"errors"
	"fmt"
)

type SimpleCache struct {
	storage map[string]interface{}
}

func NewSimpleCache() *SimpleCache {
	return &SimpleCache{
		storage: make(map[string]interface{}),
	}
}

func (this *SimpleCache) generateKey(service, key string) string {
	return fmt.Sprintf("%s.%s", service, key)
}

func (this *SimpleCache) SetCacheValue(service, key string, value interface{}, expire int) error {
	k := this.generateKey(service, key)
	this.storage[k] = value
	return nil
}

func (this *SimpleCache) GetCacheValue(service, key string, vtype interface{}) (interface{}, error) {
	k := this.generateKey(service, key)
	v, ok := this.storage[k]
	if !ok {
		return nil, errors.New("key " + k + " is not exist")
	}
	return v, nil
}

func (this *SimpleCache) HasCacheKey(service, key string) bool {
	k := this.generateKey(service, key)
	_, ok := this.storage[k]
	return ok
}

func (this *SimpleCache) ResetCacheKey(service, key string) error {
	k := this.generateKey(service, key)
	delete(this.storage, k)
	return nil
}
