package main

import "fmt"

func main() {
	p0 := new(int)   // p0 points to a zero int value.
	fmt.Println(p0)  // (a hex address string)
	fmt.Println(*p0) // 0

	// x is a copy of the value at
	// the address stored in p0.
	x := *p0
	// Both take the address of x.
	// x, *p1 and *p2 represent the same value.
	p1, p2 := &x, &x
	fmt.Println(p1 == p2) // true
	fmt.Println(p0 == p1) // false
	p3 := &*p0            // <=> p3 := &(*p0) <=> p3 := p0
	// Now, p3 and p0 store the same address.
	fmt.Println(p0 == p3) // true
	*p1, *p2 = 123, 789
	fmt.Println(*p1, *p2, x, *p3) // 789 789 123
	fmt.Println(&x, &p1, &p2)
}
