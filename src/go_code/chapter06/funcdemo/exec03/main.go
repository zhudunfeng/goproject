package main

import (
	"fmt"
)

func main() {
	a := 10
	b := 20
	// swap(&a, &b)
	// swap2(swap, &a, &b)
	f := swap3()
	f(&a, &b)
	fmt.Printf("a=%d,b=%d\n", a, b)
}

func swap(x *int, y *int) {
	*x = *x ^ *y
	*y = *x ^ *y
	*x = *x ^ *y
}

func swap2(myFunc func(*int, *int), x *int, y *int) {
	myFunc(x, y)
}

func swap3() func(*int, *int) {
	return func(x *int, y *int) {
		tmp := x
		x = y
		y = tmp
	}
}
