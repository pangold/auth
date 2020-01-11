package system

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

type RedisCache struct {
	conn redis.Conn
}

func NewRedisCache(host string) *RedisCache {
	conn, err := redis.Dial("tcp", host)
	if err != nil {
		panic(err.Error())
	}
	return &RedisCache{
		conn: conn,
	}
}

func (this *RedisCache) generateKey(service, key string) string {
	return fmt.Sprintf("%s.%s", service, key)
}

func (this *RedisCache) HasCacheKey(service, key string) bool {
	k := this.generateKey(service, key)
	exist, err := redis.Bool(this.conn.Do("EXISTS", k))
	if err != nil {
		return false
	}
	return exist
}

func (this *RedisCache) SetCacheValue(service, key string, value interface{}, expire int) error {
	k := this.generateKey(service, key)
	if _, err := this.conn.Do("SET", k, value); err != nil {
		return err
	}
	if expire < 0 {
		return nil
	}
	if _, err := this.conn.Do("EXPIRE", k, expire); err != nil {
		this.conn.Do("DEL", key)
		return err
	}
	return nil
}

func (this *RedisCache) GetCacheValue2(service, key string) (interface{}, error) {
	k := this.generateKey(service, key)
	// FIXME: Is there a better way? to fit any types
	v, err := redis.String(this.conn.Do("GET", k))
	if err != nil {
		return nil, err
	}
	return v, nil
}

// not perfect. due to the vtype...
func (this *RedisCache) GetCacheValue(service, key string, vtype interface{}) (interface{}, error) {
	k := this.generateKey(service, key)
	switch vtype.(type) {
	case bool:
		return redis.Bool(this.conn.Do("GET", k))
	case []byte:
		return redis.Bytes(this.conn.Do("GET", k))
	case string:
		return redis.String(this.conn.Do("GET", k))
	case []string:
		return redis.Strings(this.conn.Do("GET", k))
	case map[string]string:
		return redis.StringMap(this.conn.Do("GET", k))
	case float64:
		return redis.Float64(this.conn.Do("GET", k))
	case []float64:
		return redis.Float64s(this.conn.Do("GET", k))
	case int:
		return redis.Int(this.conn.Do("GET", k))
	case []int:
		return redis.Ints(this.conn.Do("GET", k))
	case map[string]int:
		return redis.IntMap(this.conn.Do("GET", k))
	case int64:
		return redis.Int64(this.conn.Do("GET", k))
	case []int64:
		return redis.Int64s(this.conn.Do("GET", k))
	case map[string]int64:
		return redis.Int64Map(this.conn.Do("GET", k))
	case uint64:
		return redis.Uint64(this.conn.Do("GET", k))
	}
	return nil, nil
}

func (this *RedisCache) ResetCacheKey(service, key string) error {
	k := this.generateKey(service, key)
	if _, err := this.conn.Do("DEL", k); err != nil {
		return err
	}
	return nil
}
