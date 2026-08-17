package cache_test

import (
	"fmt"

	"github.com/yuvv/candy/cache"
)

func ExampleLRU() {
	lru := cache.NewLRU[string, int](2)
	lru.Put("a", 1)
	lru.Put("b", 2)

	value, found := lru.Get("a") // Promote "a" to most recently used.
	fmt.Println(value, found)

	lru.Put("c", 3) // Evicts "b".
	_, bFound := lru.Get("b")
	cValue, cFound := lru.Get("c")
	fmt.Println(bFound, cValue, cFound)
	// Output:
	// 1 true
	// false 3 true
}
