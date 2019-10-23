package cache


type Cache interface {
	HasCacheKey(service, key string) bool
	SetCacheValue(service, key string, value interface{}, expire int) error
	GetCacheValue(service, key string, vtype interface{}) (interface{}, error)
	ResetCache(service, key string) error
}
