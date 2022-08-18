package main

import "fmt"

type myFunc func(int, int) int

func getSum(num1, num2 int) int {
	return num1 + num2
}

func sum(getSum myFunc, num1 int, num2 int) int {
	return getSum(num1, num2)
}

func main() {
	num1 := 10
	num2 := 20
	res := sum(getSum, num1, num2)
	fmt.Println("res=", res)
}
