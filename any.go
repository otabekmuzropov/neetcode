package main

import (
	"fmt"
)

func main() {
	a("")
}

func a(a string) {
	var b *string = &a
	fmt.Println(b)
}
