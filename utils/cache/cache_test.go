package cache

import (
	"testing"
)

var (
	c Cache
)

func init() {
	c = UseSimpleCache()
	//c = UseRedigoCache("127.0.0.1:6379")
	if c == nil {
		panic("invalid cache handler")
	}
}

func TestSetCacheValue(t *testing.T) {
	n, s := int(10), "abc"
	if err := c.SetCacheValue("unittest", "key1", n, 60); err != nil {
		t.Errorf("failed to set cache value: %d", n)
		return
	}
	if err := c.SetCacheValue("unittest", "key2", s, 10); err != nil {
		t.Errorf("falied to set cache value: %s", s)
		return
	}
}

func TestGetCacheValue(t *testing.T) {
	n, s := int(10), "abc"
	if v, err := c.GetCacheValue("unittest", "key1", n); err != nil || v == nil || v.(int) != n {
		t.Errorf("failed to get cache value, error: %v, %v", err, v)
		return
	}
	if v, err := c.GetCacheValue("unittest", "key2", s); err != nil || v == nil || v.(string) != s {
		t.Errorf("failed to get cache value, error: %v, %v", err, v)
		return
	}

}

func TestHasCacheValue1(t *testing.T) {
	if !c.HasCacheKey("unittest", "key1") || !c.HasCacheKey("unittest", "key2") {
		t.Errorf("error: both key1 and key2 should be exist")
	}
}

func TestResetCache(t *testing.T) {
	if err := c.ResetCache("unittest", "key1"); err != nil {
		t.Errorf("error: failed to remove key %s", "key1")
	}
}

func TestHasCacheValue2(t *testing.T) {
	if c.HasCacheKey("unittest", "key1") {
		t.Errorf("error: key1 should not be exist")
	}
}

