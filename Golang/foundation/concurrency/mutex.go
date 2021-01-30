package concurrency

import (
	"fmt"
	"sync"
)

func RunMutex() {
	c := SafeCounter{stringToCount: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		c.Increment("key1")
	}
	fmt.Println(c.Value("key1"))
}

type SafeCounter struct {
	mu sync.Mutex
	stringToCount map[string]int
}

func (counter *SafeCounter) Increment(key string) {
	counter.mu.Lock()
	defer counter.mu.Unlock()
	counter.stringToCount[key]++
}

func (counter *SafeCounter) Value(key string) int{
	counter.mu.Lock()
	defer counter.mu.Unlock()
	return counter.stringToCount[key]
}