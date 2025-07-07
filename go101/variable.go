package main

import "fmt"

func main() {
	fmt.Println(int(^uint(0) >> 1))
	fmt.Println(uint(0))
	fmt.Println(^uint(0))
	fmt.Println(int(^uint(0) >> 1))
}
