package main

import (
	"fmt"
	"sync"
)

func main() {
	var memoryAccess sync.Mutex
	var val int
	go func() {
		memoryAccess.Lock()
		val++
		memoryAccess.Unlock()
	}()

	memoryAccess.Lock()
	if val == 0 {
		fmt.Printf("the value is %v. \n", val)
	} else {
		fmt.Printf("the value is %v.\n", val)
	}

	memoryAccess.Unlock()

}
