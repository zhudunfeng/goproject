package main

import "fmt"

func main() {
	var slice []int = make([]int, 4, 10)
	//对于切片，必须make使用
	fmt.Println(slice) //默认值为0
	fmt.Println("slice len=", len(slice), "slice cap=", cap(slice))
	slice[0] = 100
	slice[2] = 200
	fmt.Println(slice)

	//方式3
	fmt.Println()
	//第3种方式:定义一个切片，直接就指定具体数组，使用原理类似make的方式。
	var strSlice []string = []string{"tom", "jack", "mary"}
	fmt.Println("strSlice=", strSlice)          //[tom jack mary]
	fmt.Println("strSlice len=", len(strSlice)) //3
	fmt.Println("strSlice cap=", cap(strSlice)) //3

}
