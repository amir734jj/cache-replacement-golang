package cache_replacement_golang

import (
	"strconv"
	"testing"
)

func validate(t *testing.T, key string, err error, expected interface{}, actual interface{}) {
	if err != nil {
		t.Errorf("Failed to find cached value %s", key)
	} else if expected != actual {
		t.Errorf("Mismatch value for a key %s", key)
	}
}

func CacheGenericTesting(t *testing.T, makeCache func(int) Cache) {
	size := 256
	cache := makeCache(size)

	// Test add and immediate get
	for i := 0; i < size; i++ {
		key := strconv.Itoa(i)
		cache.Add(key, i)
		value, err := cache.Get(key)
		validate(t, key, err, i, value)
	}

	// Test get afterward
	for i := 0; i < size; i++ {
		key := strconv.Itoa(i)
		value, err := cache.Get(key)
		validate(t, key, err, i, value)
	}

	// Test update
	for i := 0; i < size; i++ {
		key := strconv.Itoa(i)
		cache.Add(key, i+i)
		value, err := cache.Get(key)
		validate(t, key, err, i+i, value)
	}
}

func TestFIFO(t *testing.T) {
	CacheGenericTesting(t, FIFO())
}

func TestFILO(t *testing.T) {
	CacheGenericTesting(t, FILO())
}

func TestLRU(t *testing.T) {
	CacheGenericTesting(t, LRU())
}

func TestMRU(t *testing.T) {
	CacheGenericTesting(t, MRU())
}

func TestLFU(t *testing.T) {
	CacheGenericTesting(t, LFU())
}
