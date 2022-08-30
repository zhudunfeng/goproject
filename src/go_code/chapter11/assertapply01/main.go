package main

import "fmt"

//声明/定义一个接口
type Usb interface {
	//声明两个没有实现的方法
	Start()
	Stop()
}

type Phone struct {
	name string
}

//定义Phone特有方法Call
func (p Phone) Call() {
	fmt.Println("手机开始打电话....")
}

//让Phone实现Usb的两方法
func (p Phone) Start() {
	fmt.Println("手机开始工作...")
}

func (p Phone) Stop() {
	fmt.Println("手机停止工作...")
}

type Camera struct {
	name string
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
	//类型断言
	if phone, ok := usb.(Phone); ok {
		phone.Call()
	}
	usb.Stop()
}

func main() {
	//定义一个Usb接口数组，可以存放Phone和Camera的结构体变量
	//这里就体现出多态数组
	var usbArr [3]Usb

	usbArr[0] = Phone{"vivo"}
	usbArr[1] = Phone{"iphone"}
	usbArr[2] = Camera{"佳能"}

	//Phone还有一个特有的方法call()，请遍历Usb数组，如果是Phone变量，
	//除了调用Usb接口声明的方法外，还需要调用Phone 特有方法 call.=》类型断言
	var computer Computer
	for _, usb := range usbArr {
		computer.WorkIng(usb)
		fmt.Println()
	}
}
