package main

import "fmt"

//声明/定义一个接口
type Usb interface {
	//声明两个没有实现的方法
	Start()
	Stop()
}

type Phone struct{}

//让Phone实现Usb的两方法
func (p Phone) Start() {
	fmt.Println("手机开始工作...")
}

func (p Phone) Stop() {
	fmt.Println("手机停止工作...")
}

type Camera struct {
}

func (c Camera) Start() {
	fmt.Println("相机开始工作...")
}

func (c Camera) Stop() {
	fmt.Println("相机停止工作...")
}

type Computer struct{}

//编写一个方法working方法,接收一个Usb接口类型变量
//只要是实现了usb接口（所谓实现usb接口，就是指实现了Usb接口声明所有方法)
func (c Computer) WorkIng(usb Usb) {
	//通过usb接口变量来调用start和stop方法
	usb.Start()
	usb.Stop()
}

func main() {
	//测试
	//先创建结构体
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}

	//关键点
	computer.WorkIng(phone)
	computer.WorkIng(camera)
}
