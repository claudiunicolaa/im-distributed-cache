package imdistributedcache_test

import (
	"strconv"
	"testing"

	imdistributedcache "github.com/claudiunicolaa/im-distributed-cache"
)

func BenchmarkCache_Get(b *testing.B) {
	c := imdistributedcache.NewCache()

	// Populate the cache with some data
	for i := 0; i < b.N; i++ {
		c.Set(strconv.Itoa(i), i, 100)
	}

	b.ResetTimer()

	// Measure the time it takes to retrieve values from the cache
	for i := 0; i < b.N; i++ {
		c.Get(strconv.Itoa(i % 1000))
	}
}
