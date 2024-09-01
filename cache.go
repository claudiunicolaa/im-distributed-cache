package imdistributedcache

import (
	"sync"
	"time"
)

// Item represents a cache item that contains the data and the expiration time
type Item struct {
	data     interface{}
	expireAt int64
}

// NewItem creates a new cache item
func NewItem(data interface{}, expireAt int64) Item {
	return Item{
		data:     data,
		expireAt: expireAt,
	}
}

// Data returns the data of the cache item
func (i Item) Data() interface{} {
	return i.data
}

// ExpireAt returns the expiration time of the cache item
func (i Item) ExpireAt() int64 {
	return i.expireAt
}

// Cache is the interface that wraps the basic cache operations
type Cache interface {
	// Get retrieves an item from the cache
	Get(key string) (Item, error)
	// Set adds an item to the cache. The ttl is the time to live in seconds.
	Set(key string, value interface{}, ttl int64) error
	// Delete removes an item from the cache
	Delete(key string) error
}

type cache struct {
	items map[string]Item
	mux   sync.Mutex
}

// NewCache creates a new cache instance with an empty storage
func NewCache() Cache {
	return &cache{
		items: make(map[string]Item),
	}
}

func (c *cache) Get(key string) (Item, error) {
	c.mux.Lock()
	defer c.mux.Unlock()

	item, ok := c.items[key]
	if !ok {
		return Item{}, ErrCacheMiss
	}

	now := time.Now().Unix()
	if item.ExpireAt() > 0 && item.ExpireAt() < now {
		delete(c.items, key)
		return Item{}, ErrCacheMiss
	}

	return item, nil
}

func (c *cache) Set(key string, value interface{}, ttl int64) error {
	// too short ttl, do not set the item
	if ttl < 1 {
		return nil
	}

	c.mux.Lock()
	defer c.mux.Unlock()

	c.items[key] = NewItem(value, time.Now().Unix()+ttl)

	return nil
}

func (c *cache) Delete(key string) error {
	c.mux.Lock()
	defer c.mux.Unlock()

	delete(c.items, key)
	return nil
}
