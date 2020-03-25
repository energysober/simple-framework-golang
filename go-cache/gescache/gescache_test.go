package gescache

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetter(t *testing.T) {
	var f Getter = GetterFunc(func(key string) ([]byte, error) {
		return []byte(key), nil
	})

	expect := []byte("key")
	if v, _ := f.Get("key"); !reflect.DeepEqual(v, expect) {
		t.Fatal("callback failed")
	}
}

var db = map[string]string{
	"Windy": "100",
	"Mack":  "99",
	"Stone": "80",
}

func TestGet(t *testing.T) {
	loadCounts := make(map[string]int, len(db))

	ges := NewGroup("scores", 2<<10, GetterFunc(func(key string) ([]byte, error) {
		if value, ok := db[key]; ok {
			if _, ok := loadCounts[key]; !ok {
				loadCounts[key] = 1
			}
			loadCounts[key] += 1
			return []byte(value), nil
		}
		return nil, fmt.Errorf("%s not exist", key)
	}))

	for k, v := range db {
		if view, err := ges.Get(k); err != nil || view.String() != v {
			t.Fatalf("failed to get value of %s", v)
		}
		if _, err := ges.Get(k); err != nil || loadCounts[k] > 1 {
			t.Fatalf("cache %s miss", k)
		}
	}

	if view, err := ges.Get("unknown"); err == nil {
		t.Fatalf("the value of unknow should be empty, but %s got", view)
	}
}
