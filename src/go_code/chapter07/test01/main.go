package main

import (
	"fmt"
)

func main() {
	//创建一个数组
	var array = [5]int{1, 2, 5, 4, 3}
	//数组冒泡排序
	arraySort(&array)
	fmt.Println(array)

	var num byte
	num = 'S'
	fmt.Printf("num=%v\n",num)
	fmt.Printf("%T\n", num);

	var s string ="S"
	fmt.Printf("s=%v\n",s)
	
	
}

func arraySort(arr *[5]int) {
	// 数组冒泡排序
	//外层循环控制轮数
	for i := 0; i < len(arr)-1; i++ {
		//内层控制循环次数
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] < arr[j+1] {
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
	}
}
