package main

import (
	"fmt"
	"runtime"
)

func main() {
	var before runtime.MemStats
	runtime.ReadMemStats(&before)

	a := [10 * 1024 * 1024]byte{}
	println(&a)

	b := [10*1024*1024 + 1]byte{}
	println(&b)

	var after runtime.MemStats
	runtime.ReadMemStats(&after)
	memoryUsed := (after.TotalAlloc - before.TotalAlloc) / 1024 / 1024
	fmt.Println("Memory used:", memoryUsed, "MB")
}
