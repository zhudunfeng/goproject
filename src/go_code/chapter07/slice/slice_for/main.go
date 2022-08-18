package main

import "fmt"

func main() {
	//数组
	var intArr [5]int = [...]int{11, 22, 33, 44, 55}
	//切片
	slice := intArr[1:4] //22, 33, 44
	//使用常规的for循环遍历切片
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slcie[%v]=%v ", i, slice[i])
	}

	fmt.Println()

	//使用for--range 方式遍历切片
	for i, v := range slice {
		fmt.Printf("i=%v v=%v\n", i, v)
	}

	//切片后仍然可以切片
	slice2 := slice[1:2] //33
	slice2[0] = 100      //因为arr , slice 和slice2 指向的数据空间是同一个，因此slice2[0]=100,
	fmt.Printf("slice2=%v\n", slice2)
	fmt.Printf("slice=%v\n", slice)
	fmt.Printf("intArr=%v\n", intArr)
}
