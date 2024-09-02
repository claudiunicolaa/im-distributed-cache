package imdistributedcache

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
