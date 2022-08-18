package main

import (
	"fmt"
)

func main() {
	var i int = 9

	//指针变量用于存放地址
	var ptr *int = &i
	var ptr1 **int = &ptr

	fmt.Printf("ptr的值%v\n", ptr)    //这是i的地址值
	fmt.Printf("ptr的地址%v\n", &ptr)  //这是ptr的地址值
	fmt.Printf("ptr指向的值%v\n", *ptr) //这是i的值


	// 二级指针
	fmt.Printf("ptr1的值%v\n", ptr1)    //这是ptr的地址
	fmt.Printf("ptr1的地址%v\n", &ptr1)  //这是ptr1的地址值
	fmt.Printf("ptr1保存的地址%v\n", *ptr1)  //这是i的地址值
	fmt.Printf("ptr1指向的值%v\n", **ptr1) //这是i的值

	var nums = []int{1, 3, 2, 1, 4}
	// ChangeNum(nums)
	fmt.Println(nums)
	i2 := DuplicateNum(nums)
	fmt.Printf("%d\n", i2)

	

}
