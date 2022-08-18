package main

import (
	"fmt"
)

func main() {
	//数组定义的四种方式
	//第一种方式
	var numArray1 [3]int = [3]int{1, 2, 3}

	//第二种方式
	var numArray2 = [3]int{4, 5, 6}

	//第三种方式
	var numArray3 = [...]int{7, 8, 9}

	//第四种方式
	//var numArray4 = [3]int{1: 900, 0: 800, 2: 1000}

	//可以使用类型推断
	numArray4 := [3]int{1: 900, 0: 800, 2: 1000}

	fmt.Println(numArray1)
	fmt.Println(numArray2)
	fmt.Println(numArray3)
	fmt.Println(numArray4)
}
