package main

import (
    "fmt"
    "runtime"
    "sync"
    "time"
)

// Simulate a time-consuming task
func task(id int) {
    fmt.Printf("Task %d started\n", id)
    time.Sleep(1 * time.Second) // Simulate some work (e.g., a network call)
    fmt.Printf("Task %d completed\n", id)
}

func main() {
    // Set number of CPUs to use
    runtime.GOMAXPROCS(2) // Set it to 2 to simulate parallelism with 2 cores. Change this to 1 to simulate concurrency with 1 core.

    var wg sync.WaitGroup
    taskCount := 5

    fmt.Println("Start of tasks:")

    // Run tasks concurrently
    for i := 1; i <= taskCount; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            task(id)
        }(i)
    }

    wg.Wait()
    fmt.Println("All tasks completed!")
}
