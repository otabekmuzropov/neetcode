package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	printsum := func(v1, v2 *value) {
		defer wg.Done()
		v1.mu.Lock()
		defer v1.mu.Unlock()

		v2.mu.Lock()
		defer v2.mu.Unlock()
		fmt.Printf("sum=%v\n", v1.value+v2.value)
	}

	var a, b value
	wg.Add(2)
	go printsum(&a, &b)
	go printsum(&b, &a)
	wg.Wait()
}

type value struct {
	mu    sync.Mutex
	value int
}
