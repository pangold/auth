package utils

import (
	"testing"
)

func TestHasCacheKey(t *testing.T) {
	if HasCacheKey("test", "key1000") {
		t.Errorf("test.key1 shouldn't be exist")
	}
}

func TestSetCacheValue(t *testing.T) {
	if err := SetCacheValue("test", "key1", int(100), 1); err != nil {
		t.Errorf("setting cache failed")
	}
}

func TestHasCacheKey2(t *testing.T) {
	if !HasCacheKey("test", "key1") {
		t.Errorf("test.key1 should be exist")
	}
}

func TestGetCacheValue(t *testing.T) {
	if v, err := GetCacheValue("test", "key1", int(0)); v == nil || err != nil {
		t.Errorf("getting cache failed, value: %v, err: %v", v, err)
	}
}

