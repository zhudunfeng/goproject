package main

import "fmt"

type Circle struct {
	radius float64
}

//2)声明一个方法area和circle绑定，可以返回面积。
//值传递
func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

//为了提高效率,通常我们方法和结构体的指针类型绑定
//引用传递
func (c *Circle) area1() float64 {
	//因为c是指针，因此我们标准的访问其字段的方式是（*c ) .radius
	// return 3.14 * (*c).radius * (*c).radius
	//(*c).radius等价c.radius
	return 3.14 * c.radius * c.radius
}
func main() {
	// 1)声明一个结构体circle,字段为radius
	//2)声明一个方法area和circle绑定，可以返回面积。
	//3)提示:画出area执行过程+说明
	// var circle Circle
	// circle.radius = 4.0
	// res := circle.area()
	// fmt.Println("面积是：", res)
	var circle Circle
	circle.radius = 4.0
	// res2 := (&circle).area1()
	//编译器底层做了优化(&c).area2(）等价c.area()
	//因为编译器会自动的给加上 &c
	res2 := circle.area1()
	fmt.Println("面积是：", res2)
}
