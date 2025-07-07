package main

import (
	"fmt"
	"time"
)

func writer(ch chan int) {
	ch <- 42 // Writing
}

func reader(ch chan int) {
	fmt.Println(<-ch) // Reading
}

func main() {
	ch := make(chan int)
	go reader(ch)
	//go writer(ch)
	time.Sleep(time.Second)
}
