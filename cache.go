package cache

import (
	"sync"
	"time"
)

type entry[K comparable, V any] struct {
	Key      K
	Val      V
	ExpireAt time.Time
}

type Cache[K comparable, V any] interface {
	Put(k K, v V, ttl time.Duration) error
	Get(k K) (V, bool)
	Del(k K) error
}

type InMemoryCache[K comparable, V any] struct {
	store     map[K]entry[K, V]
	storeLock sync.Mutex
}

func NewInMemoryCache[K comparable, V any]() Cache[K, V] {
	return &InMemoryCache[K, V]{
		store: make(map[K]entry[K, V]),
	}
}

func (c *InMemoryCache[K, V]) Put(k K, v V, ttl time.Duration) error {
	c.storeLock.Lock()
	defer c.storeLock.Unlock()
	c.store[k] = entry[K, V]{k, v, time.Now().Add(ttl)}
	return nil
}

func (c *InMemoryCache[K, V]) Get(k K) (V, bool) {
	c.storeLock.Lock()
	defer c.storeLock.Unlock()

	v, ok := c.store[k]
	if ok && v.ExpireAt.Sub(time.Now()) < 0 {
		return v.Val, false
	}
	return v.Val, ok
}

func (c *InMemoryCache[K, V]) Del(k K) error {
	c.storeLock.Lock()
	defer c.storeLock.Unlock()

	delete(c.store, k)
	return nil
}
