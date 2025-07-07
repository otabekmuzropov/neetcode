package main

import "fmt"

func changeMap(m1 map[string]int) {
	m1 = map[string]int{"hello": 2}
	fmt.Println(&m1)
}

func main() {
	m1 := map[string]int{"hello": 1}
	fmt.Println(m1)
	changeMap(m1)
	println(m1["hello"]) // 1
}
