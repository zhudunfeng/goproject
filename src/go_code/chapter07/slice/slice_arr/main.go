package main

import "fmt"

func main() {
	//演示切片的基本使用
	var intArr [5]int = [...]int{1, 22, 33, 66, 99}
	//声明/定义一个切片
	//slice := intArr[1:3]
	//1. slice就是切片名
	//2. intArr[1:3]表示slice引用到intArr这个数组
	//3．引用intArr数组的起始下标为1,最后的下标为3(但是不包含3)
	slice := intArr[1:3]
	fmt.Println("intArr=", intArr)
	fmt.Println("slice的元素是：", slice)        //[22 33]
	fmt.Println("slice的元素个数是：", len(slice)) //2
	fmt.Println("slice的容量是：", cap(slice))   // 4

	fmt.Printf("intArr[1]地址是%p\n", &intArr[1]) //0xc00000a4e8
	fmt.Printf("scice[0]地址是%p\n", &slice[0])   //0xc0000cc03

}
