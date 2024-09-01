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
		// compute based on the index a ttl between 0 and 30 the distribution deterministic
		// with 10% of the items having a ttl of 0
		ttl := int64(i % 30)
		if ttl == 0 {
			ttl = 1
		}
		c.Set(strconv.Itoa(i), i, ttl)
	}

	b.ResetTimer()

	// Measure the time it takes to retrieve values from the cache in same order as they were set
	for i := 0; i < b.N; i++ {
		c.Get(strconv.Itoa(i % 1000))
	}
}
