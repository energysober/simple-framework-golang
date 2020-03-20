package lru

import (
	"reflect"
	"testing"
)

// String
type String string

func (d String) Len() int {
	return len(d)
}

func TestGet(t *testing.T) {
	lru := New(int64(0), nil)
	lru.Add("testKey1", String("testValue1"))
	if v, ok := lru.Get("testKey1"); !ok || v != String("testValue1") {
		t.Fatal("cache hit testKey1=testValue1 failed")
	}
	if _, ok := lru.Get("testKey2"); ok {
		t.Fatal("cache miss testKey2 failed")
	}
}

func TestRemoveOldest(t *testing.T) {
	k1, k2, k3 := "k1", "k2", "k3"
	v1, v2, v3 := "v1", "v2", "v3"
	cacheCap := len(k1 + v1 + k2 + v2)
	lru := New(int64(cacheCap), nil)
	lru.Add(k1, String(v1))
	lru.Add(k2, String(v2))
	lru.Add(k3, String(v3))

	if _, ok := lru.Get(v1); ok || lru.Len() != 2 {
		t.Fatal("cache remove oldest k1 failed")
	}
}

func TestOnEvicted(t *testing.T) {
	keys := make([]string, 0)
	callback := func(key string, value Value) {
		keys = append(keys, key)
	}

	lru := New(int64(10), callback)
	lru.Add("key1", String("123456"))
	lru.Add("k2", String("k2"))
	lru.Add("k3", String("k3"))
	lru.Add("k4", String("k4"))

	expect := []string{"key1", "k2"}

	if !reflect.DeepEqual(expect, keys) {
		t.Fatalf("Call OnEvicted failed, expect keys equals to %s", expect)
	}
}
