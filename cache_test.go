package cache

import (
	"testing"
	"time"
)

func TestInMemoryCache(t *testing.T) {
	t.Run("test put/get/del", func(t *testing.T) {
		mem := NewInMemoryCache[string, int]()

		err := mem.Put("a", 1, time.Second)
		if err != nil {
			t.Error("Put operation should not end up with error")
		}

		v, ok := mem.Get("a")
		if !ok {
			t.Error("key 'a' does not exists in cache")
		}

		if v != 1 {
			t.Errorf("wrong value: 1 != %d", v)
		}

		v, ok = mem.Get("b")
		if ok {
			t.Error("key 'b' is not supposed to be in cache")
		}

		err = mem.Del("a")
		if err != nil {
			t.Error("Del operation should not end up with error")
		}

		v, ok = mem.Get("a")
		if ok {
			t.Error("key 'a' should not exists in cache")
		}
	})

	t.Run("test key ttl", func(t *testing.T) {
		mem := NewInMemoryCache[string, int]()

		err := mem.Put("a", 1, 500*time.Millisecond)
		if err != nil {
			t.Error("Put operation should not end up with error")
		}

		v, ok := mem.Get("a")
		if !ok {
			t.Error("key 'a' does not exists in cache")
		}

		if v != 1 {
			t.Errorf("wrong value: 1 != %d", v)
		}

		time.Sleep(501 * time.Millisecond)
		v, ok = mem.Get("a")
		if ok {
			t.Error("key 'a' supposed to be expired")
		}
	})
}
