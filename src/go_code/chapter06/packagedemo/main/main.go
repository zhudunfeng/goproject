package main

import (
	"fmt"
	//起别名
	// util "go_code/chapter06/packagedemo/utils"
	"go_code/chapter06/packagedemo/utils"
)

func main() {

	var n1 float64 = 1.2
	var n2 float64 = 1.3
	var operator byte = '+'
	r := utils.Cal(n1, n2, operator)
	fmt.Println("结果是", r)

}
