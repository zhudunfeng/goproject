package main

import "fmt"

func main() {
	//切片的拷贝操作
	//切片使用copy内置函数完成拷贝,举例说明:
	var slice4 []int = []int{1, 2, 3, 4, 5}
	var slice5 []int = make([]int, 10)

	copy(slice5, slice4)

	fmt.Println("slice4", slice4) //[1 2 3 4 5]
	fmt.Println("slice5", slice5) //[1 2 3 4 5 0 0 0 0 0]
}
