package main

import "fmt"

//猴子吃桃
func peach(n int) int {

	if n < 1 || n > 10 {
		fmt.Println("超出最大赋值范围")
		return 0
	}
	if n == 10 {
		return 1
	} else {
		return (peach(n+1) + 1) * 2
	}
}

func main() {
	fmt.Println("第1天吃了", peach(1))
	fmt.Println("第9天吃了", peach(9))
}
