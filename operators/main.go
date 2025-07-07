package main

import (
	"fmt"
	"reflect"
)

type A bool
type B = bool

func main() {
	var a, b uint8 = 255, 3
	// Compiles ok, higher overflowed bits are truncated.
	var c = a + b // c == 0
	// Compiles ok, higher overflowed bits are truncated.
	var d = a << b // d == 254
	fmt.Printf("%b\n", 2040)
	//fmt.Println(strconv.FormatInt(510, 2))
	fmt.Println("sakfdksdf", 0b11111111000)
	fmt.Println(0x1FFFFFFFF)
	fmt.Println(a, c, d)

	var aa *A
	var aaa any
	aaa = aa

	var bb *B
	var bbb any
	bbb = bb

	fmt.Println(aaa == nil)
	fmt.Println(bbb == nil)
	fmt.Println(aaa == bbb)

	const integer = 15
	const r = 'A'
	const name = 1.2
	const s = 1i
	var res = integer + r + name + s
	fmt.Println(res, reflect.TypeOf(s))
	fmt.Println("0x7FFFFFFF", 0x7FFFFFFF)

	fmt.Println("0x63", 0x63)

	var a1, b1 uint8 = 255, 1
	var d1 = a1 << b1
	fmt.Println("d is ", d1)

	fmt.Println(reflect.TypeOf(d1))
	fmt.Println(3 / 2 * 0.1)

}
