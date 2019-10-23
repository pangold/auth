package cache

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

type RedigoCache struct {
	conn redis.Conn
}

func UseRedigoCacheWithError(hostName string) (Cache, error) {
	fmt.Println("use my token")
	c, err := redis.Dial("tcp", hostName)
	if err != nil {
		return nil, err
	}
	// FIXME: conn.Close()
	return RedigoCache{conn: c}, nil
}

func UseRedigoCache(hostName string) Cache {
	token, err := UseRedigoCacheWithError(hostName)
	if err != nil {
		return nil
	}
	return token
}

func (rc RedigoCache) HasCacheKey(service, key string) bool {
	k := generateCacheKey(service, key)
	exist, err := redis.Bool(rc.conn.Do("EXISTS", k))
	if err != nil {
		return false
	}
	return exist
}

func (rc RedigoCache) SetCacheValue(service, key string, value interface{}, expire int) error {
	k := generateCacheKey(service, key)
	if _, err := rc.conn.Do("SET", k, value); err != nil {
		return err
	}
	if expire < 0 {
		return nil
	}
	if _, err := rc.conn.Do("EXPIRE", k, expire); err != nil {
		rc.conn.Do("DEL", key)
		return err
	}
	return nil
}

func (rc RedigoCache) GetCacheValue2(service, key string) (interface{}, error) {
	k := generateCacheKey(service, key)
	// FIXME: Is there a better way? to fit any types
	v, err := redis.String(rc.conn.Do("GET", k))
	if err != nil {
		return nil, err
	}
	return v, nil
}

// not perfect. due to the vtype...
func (rc RedigoCache) GetCacheValue(service, key string, vtype interface{}) (interface{}, error) {
	k := generateCacheKey(service, key)
	switch vtype.(type) {
	case bool:
		return redis.Bool(rc.conn.Do("GET", k))
	case []byte:
		return redis.Bytes(rc.conn.Do("GET", k))
	case string:
		return redis.String(rc.conn.Do("GET", k))
	case []string:
		return redis.Strings(rc.conn.Do("GET", k))
	case map[string]string:
		return redis.StringMap(rc.conn.Do("GET", k))
	case float64:
		return redis.Float64(rc.conn.Do("GET", k))
	case []float64:
		return redis.Float64s(rc.conn.Do("GET", k))
	case int:
		return redis.Int(rc.conn.Do("GET", k))
	case []int:
		return redis.Ints(rc.conn.Do("GET", k))
	case map[string]int:
		return redis.IntMap(rc.conn.Do("GET", k))
	case int64:
		return redis.Int64(rc.conn.Do("GET", k))
	case []int64:
		return redis.Int64s(rc.conn.Do("GET", k))
	case map[string]int64:
		return redis.Int64Map(rc.conn.Do("GET", k))
	case uint64:
		return redis.Uint64(rc.conn.Do("GET", k))
	}
	return nil, nil
}

func (rc RedigoCache) ResetCache(service, key string) error {
	k := generateCacheKey(service, key)
	if _, err := rc.conn.Do("DEL", k); err != nil {
		return err
	}
	return nil
}
