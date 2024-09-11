package pokecache

import (
	"fmt"
	"testing"
	"time"
)


func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	
	cases := [] struct {
		key string
		val []byte
	}{
		{
			key: "regularstring",
			val: []byte("data"),
		},
		{
			key: "https://www.url.com/",
			val: []byte("even more data"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			val, ok := cache.Get(c.key)
			if ok {
				t.Errorf("New Cache is not empty")
				return
			}
			cache.Add(c.key, c.val)
			val, ok = cache.Get(c.key)
			if !ok {
				t.Errorf("added key not found")
				return
			}

			if string(val) != string(c.val) {
				t.Errorf("Incorrect value returned")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const interval = 5 * time.Millisecond
	const waitTime = 5 * interval
	cache := NewCache(interval)

	key := "http://example.com"
	val := []byte("testData")

	cache.Add(key, val)

	_, ok := cache.Get(key)
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get(key)

	if ok {
		t.Errorf("Expected key to be deleted")
		return
	}

}
