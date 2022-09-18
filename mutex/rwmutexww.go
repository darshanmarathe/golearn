// :glot
package main

import (
	"fmt"
	"sync"
	"time"
)

// :show start
var cache map[int]int
var mu sync.RWMutex

func expensiveOperation(n int) int {
	// in real code this operation would be very expensive
	return n * n
}

func getCached(n int) int {
	mu.RLock()
	v, isCached := cache[n]
	mu.RUnlock()
	if isCached {
		return v
	}

	v = expensiveOperation(n)

	mu.Lock()
	cache[n] = v
	mu.Unlock()
	return v
}

func accessCache(max int) {
	total := 0
	for i := 0; i < max; i++ {
		n := getCached(i)
		total += n
	}
	fmt.Printf("total: %d\n", total)
}

// :show end

func main() {
	// :show start
	cache = make(map[int]int)

	go accessCache(5)
	go accessCache(3)
	go accessCache(2)
	accessCache(10)
	// :show end

	// for simplicity of the example
	// don't use time.Sleep() to coordinate goroutines
	// in production code
	time.Sleep(100 * time.Millisecond)
}
