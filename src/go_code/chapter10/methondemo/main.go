package main

import "fmt"

type Person struct {
	Name string
}

//给Person结构体添加speak方法,输出xxx是一个好人
func (p Person) speak() {
	fmt.Println(p.Name, "是一个BIGMAN")
}

//给Person结构体添加jisuan方法,可以计算从1+..+100o的结果，
//说明方法体内可以函数一样,进行各种运算
func (p Person) jisuan() {
	res := 0
	for i := 1; i <= 100; i++ {
		res += i
	}
	fmt.Println(p.Name, "计算结果是", res)
}

//给Person结构体jisuan2方法.该方法可以接收一个参数n，计算从1+…+n 的结果
func (p Person) jisuan2(n int) {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	fmt.Println(p.Name, "计算结果是", res)
}

//给Person结构体添加getSum方法,可以计算两个数的和，并返回结果。
func (p Person) getSum(n1 int, n2 int) int {
	return n1 + n2
}

//给Person类型绑定方法
func (p Person) test() {
	fmt.Println("test() name=", p.Name)
}

func main() {
	var p Person
	p.Name = "tom"
	p.test() //调用方法
	p.Name = "ADUN"

	//调用方法
	p.speak()
	p.jisuan()
	p.jisuan2(10)
	n1 := 10
	n2 := 20
	res := p.getSum(n1, n2)
	fmt.Println("res=", res)
}
