package main

import "fmt"

type integer int

func (i integer) print() {
	fmt.Println("i=", i)
}

//编写一个方法，改变i的值
func (i *integer) change() {
	*i = *i + 1
}

type Student struct {
	Name string
	Age  int
}

//给*Student实现方法string()
func (stu *Student) String() string {
	str := fmt.Sprintf("Name=[%v]  Age=[%v]\n", stu.Name, stu.Age)
	return str
}

func main() {
	var i integer = 10
	i.print()
	i.change()
	fmt.Println("main i =", i)

	//定义一个Student变量
	stu := Student{
		Name: "tom",
		Age:  18,
	}
	//如果你实现了*Student类型的string方法,就会自动调用
	fmt.Println(&stu)

}
