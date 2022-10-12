package main

import (
	"fmt"
	"reflect"
)

type Cal struct {
	Num1 int
	Num2 int
}

func (c Cal) GetSub(name string) {
	fmt.Printf("%s 完成了减法运行，%d - %d =%d", name, c.Num1, c.Num2, c.Num1-c.Num2)
}

// 1)编写一个Cal结构体，有两个字段Num1,和Num2。
// 2)方法GetSub(name string)
// 3)使用反射遍历cl结构体所有的字段信息.
// 4)使用反射机制完成对GetSub的调用，输出形式为
// "tom完成了减法运行，8-3=5"
func TestReflectFunc(b interface{}) {
	rType := reflect.TypeOf(b)
	rVal := reflect.ValueOf(b)
	fmt.Println(rType, rVal)

	fieldNum := rVal.Elem().NumField()
	for i := 0; i < fieldNum; i++ {
		//字段名
		fmt.Printf("fieldType[%v]=%v\n", i, rType.Elem().Field(i).Name)
		//字段类型【两种方式】
		fmt.Printf("fieldType[%v]=%v\n", i, rType.Elem().Field(i).Type)
		fmt.Printf("field[%v]=%v\n", i, rVal.Elem().Field(i).Type().Name())
	}

	methodNum := rVal.NumMethod()
	for i := 0; i < methodNum; i++ {
		//方法名
		fmt.Printf("method[%v]=%v\n", i, rType.Elem().Method(i).Name)
		//方法类型
		fmt.Printf("method[%v]=%v\n", i, rType.Elem().Method(i).Type)
		//返回值类型
		fmt.Printf("method[%v]=%v\n", i, rVal.Elem().Method(i).Type().Name())
	}
	arg := reflect.ValueOf("Tom")
	var args []reflect.Value
	args = append(args, arg)
	rVal.Elem().Method(0).Call(args)

}

func main() {
	cal := Cal{
		Num1: 10,
		Num2: 5,
	}

	TestReflectFunc(&cal)
}
