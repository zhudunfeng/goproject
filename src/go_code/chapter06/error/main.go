package main

import (
	"errors"
	"fmt"
	_ "time"
)

func test() {
	//使用defer + recover来捕获和处理异常
	//匿名函数定义后立马调用执行
	defer func() {
		// err := recover() //recover()内置函数,可以捕获到异常
		if err := recover(); err != nil {
			//说明捕获到错误
			fmt.Println("err=", err)
			//这里就可以将错误信息发送给管理员.
			fmt.Println("错误信息发送给管理员>>>", err)
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res=", res)
}

//函数去读取以配置文件init.conf的信息
//如果文件名传入不正确,我们就返回一个自定义的错误
func readConfig(name string) (err error) {
	if name == "config.ini" {
		//读取...
		return nil
	} else {
		//返回一个自定义错误
		return errors.New("读取文件错误！")
	}
}

func test02() {
	err := readConfig("config1.ini")
	if err != nil {
		//如果读取文件发生错误,就输出这个错误,并终止程序
		panic(err)
	}
	fmt.Println("test02继续执行")
}

func main() {
	//测试
	// test()
	// for {
	// 	fmt.Println("main中test()下面的代码")
	// }
	test02()
}
