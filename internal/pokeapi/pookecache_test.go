package pokeapi

import (
	"testing"
	"time"
)

const Interval = time.Millisecond * 10

func TestCreateCache(t *testing.T) {
	cache := NewCache(Interval)
	if cache.cache == nil {
		t.Error("Cache is Nil")
	}
}

func TestAddGetCache(t *testing.T) {
	cache := NewCache(Interval)

	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{
			inputKey: "key1",
			inputVal: []byte("val1"),
		},
		{
			inputKey: "key0",
			inputVal: []byte("val0"),
		},
	}

	for _, cs := range cases {
		cache.Add(cs.inputKey, cs.inputVal)
		actual, ok := cache.Get(cs.inputKey)
		if !ok {
			t.Error("Key not found")
		}
		if string(actual) != string(cs.inputVal) {
			t.Errorf("Value Does not match: %s != %s", string(actual), string(cs.inputVal))
		}
	}
}

func TestReap(t *testing.T) {
	cache := NewCache(Interval)

	key1 := "key1"
	cache.Add(key1, []byte("val1"))
	time.Sleep(Interval + time.Millisecond)
	_, ok := cache.Get(key1)
	if ok {
		t.Errorf("%s should have been reaped", key1)
	}
}

func TestReapFail(t *testing.T) {
	cache := NewCache(Interval)

	key1 := "key1"
	cache.Add(key1, []byte("val1"))
	time.Sleep(Interval / 2)
	_, ok := cache.Get(key1)
	if !ok {
		t.Errorf("%s should have been reaped", key1)
	}
}
