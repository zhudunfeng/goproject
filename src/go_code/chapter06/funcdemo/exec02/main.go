package main

import "fmt"

func swap(num1 *int, num2 *int) {
	tmp := *num1
	*num1 = *num2
	*num2 = tmp
}

func main() {
	//交换数
	a := 10
	b := 20
	swap(&a, &b)
	fmt.Println("a=", a, "b=", b)
}
