package pokecache

import (
    "fmt"
    "testing"
    "time"
)

func TestAddGet(t *testing.T) {
    const interval = 5 * time.Second
    cases := []struct {
        key string
        val []byte
    }{
        {
            key: "https://example.com",
            val: []byte("testdata"),
        },
        {
            key: "https://example.com/path",
            val: []byte("testdata2"),
        },
    }

    for i, c := range cases {
        t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
            cache := NewCache(interval)
            cache.Add(c.key, c.val)
            val, ok := cache.Get(c.key)
            if !ok {
                t.Errorf("Expected key %v to exist in cache", c.key)
                return
            }
            if string(val) != string(c.val) {
                t.Errorf("Expected value %v, got %v", c.val, val)
                return
            }
        })
    }
}

func TestReapLoop(t *testing.T) {
    const baseTime = 5 * time.Millisecond
    const waitTime = baseTime + 5*time.Millisecond
    cache := NewCache(baseTime)
    cache.Add("https://example.com", []byte("testdata"))

    _, ok := cache.Get("https://example.com")
    if !ok {
        t.Errorf("Expected key to exist in cache")
        return
    }

    time.Sleep(waitTime)
    _, ok = cache.Get("https://example.com")
    if ok {
        t.Errorf("Expected key to not exist in cache")
        return
    }
}
