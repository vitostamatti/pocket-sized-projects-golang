package cache

import (
	"slices"
	"sync"
	"time"
)

type entryWithTimeout[V any] struct {
	value   V
	expires time.Time
}

type Cache[K comparable, V any] struct {
	ttl time.Duration

	mu   sync.RWMutex
	data map[K]entryWithTimeout[V]

	maxSize           int
	chronologicalKeys []K
}

func New[K comparable, V any](maxSize int, ttl time.Duration) Cache[K, V] {
	return Cache[K, V]{
		ttl:               ttl,
		data:              make(map[K]entryWithTimeout[V]),
		maxSize:           maxSize,
		chronologicalKeys: make([]K, 0, maxSize),
	}
}

func (c *Cache[K, V]) Read(key K) (V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	var defaultValue V
	e, ok := c.data[key]
	switch {
	case !ok:
		return defaultValue, false
	case e.expires.Before(time.Now()):
		delete(c.data, key)
		return defaultValue, false
	default:
		return e.value, true
	}
}

func (c *Cache[K, V]) Upsert(key K, value V) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	_, alreadyExists := c.data[key]

	switch {
	case alreadyExists:
		c.deleteKeyValue(key)
	case len(c.data) == c.maxSize:
		c.deleteKeyValue(c.chronologicalKeys[0])
	}

	c.addKeyValue(key, value)

	return nil
}

func (c *Cache[K, V]) Delete(key K) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, key)
}

func (c *Cache[K, V]) addKeyValue(key K, value V) {
	c.data[key] = entryWithTimeout[V]{
		value:   value,
		expires: time.Now().Add(c.ttl),
	}
	c.chronologicalKeys = append(c.chronologicalKeys, key)
}

func (c *Cache[K, V]) deleteKeyValue(key K) {
	c.chronologicalKeys = slices.DeleteFunc(
		c.chronologicalKeys,
		func(k K) bool { return k == key })
	delete(c.data, key)
}
