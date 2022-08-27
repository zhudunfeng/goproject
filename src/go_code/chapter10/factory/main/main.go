package main

import (
	"fmt"
	"go_code/chapter10/factory/model"
)

func main() {
	// var stu = model.Student{
	// 	Name: "Tom",
	// 	Age:  18,
	// }

	var stu = model.NewStudent("Tom~", 99.8)
	fmt.Println(*stu)                                          //不加*会打印&{...}	指针格式
	fmt.Printf("name=%v score=%v\n", stu.Name, stu.GetScore()) //
	
}
