package main

import "fmt"

//声明接口
type BirdAble interface {
	Flying()
}

type FishAble interface {
	Swimming()
}

//Monkey结构体
type Monkey struct {
	Name string
}

func (this *Monkey) climbing() {
	fmt.Println(this.Name, "生来会爬树")
}

//LittleMonkey结构体
type LittleMonkey struct {
	Monkey
}

//让LittleMonkey实现BirdAble接口
func (this *LittleMonkey) Flying() {
	fmt.Println(this.Name, "通过学习会飞翔")
}

//让LittleMonkey实现FishAble接口
func (this *LittleMonkey) Swimming() {
	fmt.Println(this.Name, "通过学习会游泳")
}

func main() {
	littleMonkey := LittleMonkey{
		Monkey{
			Name: "孙悟空",
		},
	}
	littleMonkey.climbing()
	littleMonkey.Flying()
	littleMonkey.Swimming()
}
