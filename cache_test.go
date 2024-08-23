package imdistributedcache_test

import (
	"testing"
	"time"

	imdistributedcache "github.com/claudiunicolaa/im-distributed-cache"
)

func TestCache_Get_and_Set(t *testing.T) {
	tests := []struct {
		name             string
		key              string
		value            interface{}
		ttl              int64
		expectedValue    interface{}
		expectedExpireAt int64
		expectErr        bool
		expectedErr      error
	}{
		{
			name:             "Valid key",
			key:              "key",
			value:            "value",
			expectedValue:    "value",
			ttl:              6,
			expectedExpireAt: time.Now().Unix() + 6,
			expectErr:        false,
			expectedErr:      nil,
		},
		{
			name:             "Not found key",
			key:              "",
			value:            nil,
			expectedValue:    nil,
			ttl:              -1,
			expectedExpireAt: 0,
			expectErr:        true,
			expectedErr:      imdistributedcache.ErrCacheMiss,
		},
		{
			name:             "Expired key",
			key:              "expiredKey",
			value:            "expiredValue",
			expectedValue:    nil,
			ttl:              -1,
			expectedExpireAt: 0,
			expectErr:        true,
			expectedErr:      imdistributedcache.ErrCacheMiss,
		},
	}

	c := imdistributedcache.NewCache()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Add an item to the cache. If the key is empty, the item will not be added
			if test.key != "" {
				c.Set(test.key, test.value, test.ttl)
			}

			// Retrieve the item from the cache
			result, err := c.Get(test.key)

			// Verify the error if expected
			if (err != nil) != test.expectErr {
				t.Errorf("unexpected error: %v", err)
			}

			// Verify the retrieved item's data
			if result.Data() != test.expectedValue {
				t.Errorf("expected data to be %v, got %v", test.expectedValue, result.Data())
			}

			// Verify the retrieved item's expiration time
			if result.ExpireAt() != test.expectedExpireAt {
				t.Errorf("expected expiration time to be %v, got %v", test.expectedExpireAt, result.ExpireAt())
			}
		})
	}
}

func TestCache_Delete(t *testing.T) {
	tests := []struct {
		name          string
		key           string
		value         interface{}
		ttl           int64
		expectedValue interface{}
		expectErr     bool
		expectedErr   error
	}{
		{
			name:          "Existing key",
			key:           "key",
			value:         "value",
			ttl:           6,
			expectedValue: nil,
			expectErr:     false,
			expectedErr:   nil,
		},
		{
			name:          "Not found key",
			key:           "",
			value:         nil,
			expectedValue: nil,
			ttl:           -1,
			expectErr:     false,
			expectedErr:   nil,
		},
	}

	c := imdistributedcache.NewCache()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Add an item to the cache. If the key is empty, the item will not be added
			if test.key != "" {
				c.Set(test.key, test.value, test.ttl)
			}

			// Delete the item from the cache
			err := c.Delete(test.key)

			// Verify the error if expected
			if (err != nil) != test.expectErr {
				t.Errorf("unexpected error: %v", err)
			}

			// Retrieve the item from the cache
			result, _ := c.Get(test.key)

			// Verify that the item is deleted
			if result.Data() != test.expectedValue {
				t.Errorf("expected data to be %v, got %v", test.expectedValue, result.Data())
			}
		})
	}
}