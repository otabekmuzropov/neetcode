package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	arr := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60}}

	fmt.Println(arr)
	t := time.Now()
	runtime.GC()

	fmt.Println(time.Since(t))
}
