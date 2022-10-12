package main

import (
	"fmt"
	"reflect"
)

//专门演示反射
func testReflect(b interface{}) {
	//通过反射获取的传入变量的 type,kind,值
	//1.先获取到 reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType:", rTyp)

	//2.获取到 reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Println("rVal:", rVal)
	n2 := 2 + rVal.Int()
	fmt.Println("n2:", n2)

	fmt.Printf("rVal=%v  rVal-type=%T\n", rVal, rVal)

	//下面我们讲rVal转成interface{}
	iV := rVal.Interface()

	//将interface{} 通过断言转成需要的类型
	num2 := iV.(int)
	fmt.Printf("num2=%v num2-type=%T\n", num2, num2)
}

//专门演示反射[对结构体的反射]
func testReflect02(b interface{}) {
	//通过反射获取的传入变量的 type,kind,值
	//1.先获取到 reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType:", rTyp)

	//2.获取到 reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Println("rVal:", rVal)

	fmt.Printf("rVal=%v  rVal-type=%T\n", rVal, rVal)

	//下面我们讲rVal转成interface{}
	iV := rVal.Interface()

	//将interface{} 通过断言转成需要的类型
	//这里,我们就简单使用了一带检测的类型断言。
	//同学们可以使用swtich的断言形式来做的更加的灵活
	stu, ok := iV.(Student)
	if ok {
		fmt.Println("name:", stu.Name)
	}
}

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	//请编写一个案例,
	//演示对(基本数据类型、interface{(}、reflect.value)进行反射的基本操作

	//1．先定义一个int
	// var b int = 100
	// testReflect(b)

	//2.定义一个Student的结构体
	stu := Student{
		Name: "tom",
		Age:  18,
	}

	testReflect02(stu)

}
