package main

import (
	"fmt"
	"time"

	imdistributedcache "github.com/claudiunicolaa/im-distributed-cache"
)

func main() {
	cache := imdistributedcache.NewCache()

	// set in cache with ttl 60 seconds
	cache.Set("key", "value", 60)

	// get from cache
	item, err := cache.Get("key")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Expected value: value; Value:", item.Data()) //

	// delete from cache
	cache.Delete("key")

	// get from cache
	item, _ = cache.Get("key")
	fmt.Println("Expected value: <nil>; Value:", item.Data()) //

	// set in cache with ttl 1 seconds
	cache.Set("key", "value", 1)

	// get from cache
	item, _ = cache.Get("key")
	fmt.Println("Expected value: value; Value:", item.Data()) //

	// wait for the item to expire
	time.Sleep(2 * time.Second)

	// get from cache
	item, _ = cache.Get("key")
	fmt.Println("Expected value: <nil>; Value:", item.Data()) //
}
