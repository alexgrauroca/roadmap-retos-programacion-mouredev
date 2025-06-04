package fibonacci

import "sync"

var (
	// as many operations will be duplicated, the cache will help on performance improvements
	cache = make(map[int]int)
	// caching has race condition when there are concurrent requests, Mutex will prevent them by locking read/write ops
	mu sync.Mutex
)

func Fibonacci(n int) int {
	// if n value is cached, then we return it without doing any extra operation
	if res, found := getCache(n); found {
		return res
	}

	// this is an easy operation with low performance impact, we don't need to put it in the cache
	if n <= 1 {
		return n
	}

	// calculating and storing the fibonacci value of n
	v := Fibonacci(n-1) + Fibonacci(n-2)
	setCache(n, v)
	return v
}

func getCache(n int) (int, bool) {
	mu.Lock()
	res, found := cache[n]
	mu.Unlock()
	return res, found
}

func setCache(n int, v int) {
	mu.Lock()
	cache[n] = v
	mu.Unlock()
}
