package imdistributedcache_test

import (
	"testing"
	"time"

	imdistributedcache "github.com/claudiunicolaa/im-distributed-cache"
)

func TestItem_Data(t *testing.T) {
	tests := []struct {
		name     string
		data     interface{}
		expireAt int64
	}{
		{
			name:     "String data",
			data:     "testData",
			expireAt: time.Now().Unix() + 60,
		},
		{
			name:     "Integer data",
			data:     12345,
			expireAt: time.Now().Unix() + 60,
		},
		{
			name:     "Boolean data",
			data:     true,
			expireAt: time.Now().Unix() + 60,
		},
		{
			name: "Custom struct data",
			data: struct {
				Field1 string
				Field2 int
			}{
				Field1: "value1",
				Field2: 42,
			},
			expireAt: time.Now().Unix() + 60,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			item := imdistributedcache.NewItem(test.data, test.expireAt)

			if item.Data() != test.data {
				t.Errorf("expected data to be %v, got %v", test.data, item.Data())
			}
		})
	}
}

func TestItem_ExpireAt(t *testing.T) {
	data := "testData"
	expireAt := time.Now().Unix() + 60
	item := imdistributedcache.NewItem(data, expireAt)

	if item.ExpireAt() != expireAt {
		t.Errorf("expected expireAt to be %v, got %v", expireAt, item.ExpireAt())
	}
}
