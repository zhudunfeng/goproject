package main

import (
	"fmt"
	"reflect"
)

func reflect01(b interface{}) {
	//获取到reflect.Value
	rType := reflect.TypeOf(b)
	rVal := reflect.ValueOf(b)
	//看看rVal的Kind是什么
	fmt.Printf("rVal kind=%v  rVal Type=%v\n ", rVal.Kind(), rType)
	//rVal
	//Elem返回v持有的接口保管的值的Value封装，或者v持有的指针指向的值的Value封装。
	rVal.Elem().SetInt(20)
}

func main() {
	var num int = 10
	reflect01(&num)
	fmt.Printf("num=%d\n", num)

	//可以这样理解rVal.Elem()
	// num:=9
	// ptr *int = &num
	// num2:=*ptr //===类似 rVal.Elem()
}
