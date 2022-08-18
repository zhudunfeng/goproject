package main

import "fmt"

func main() {
	fmt.Println("开始--------------------------------")

	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i := 0; i < len(array); i++ {
		fmt.Printf("arr[%d]=%d\n", i, array[i])
	}

	fmt.Println("结束--------------------------------")

	for i := 0; i < len(array); i++ {
		fmt.Printf("arr[%d]=%d\n", i, array[i])
	}
}
