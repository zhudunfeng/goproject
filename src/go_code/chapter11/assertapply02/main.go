package main

import "fmt"

type Student struct{}

//编写一个函数,可以判断输入的参数是什么类型
func TypeJudge(items ...interface{}) {
	for index, x := range items {
		index += 1
		switch x.(type) {
		case bool:
			fmt.Printf("第%d个参数是bool类型，值是%v\n", index, x)
		case float32:
			fmt.Printf("第%d个参数是float32类型，值是%v\n", index, x)
		case float64:
			fmt.Printf("第%d个参数是float64类型，值是%v\n", index, x)
		case int, int32, int64:
			fmt.Printf("第%d个参数是整数类型，值是%v\n", index, x)
		case string:
			fmt.Printf("第%d个参数是string类型，值是%v\n", index, x)
		case byte:
			fmt.Printf("第%d个参数是byte类型，值是%c\n", index, x)
		case Student:
			fmt.Printf("第%d个参数是Student类型，值是%v\n", index, x)
		case *Student:
			fmt.Printf("第%d个参数是*Student类型，值是%v\n", index, x)
		default:
			fmt.Printf("第%d个参数是类型不确定，值是%v\n", index, x)
		}
	}
}

func main() {
	var n1 float32 = 1.1
	var n2 float64 = 2.2
	var n3 int32 = 30
	var name string = "tom"
	address := "广州"
	n4 := 300
	n5 := byte('c')

	student1 := Student{}
	student2 := &Student{}

	TypeJudge(n1, n2, n3, name, address, n4, n5, student1, student2)
}
