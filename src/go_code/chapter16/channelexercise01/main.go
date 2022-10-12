package main

import "fmt"

type Cat struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	//定义一个存放任意数据类型的管道3个数据
	// var allchan chan interface{}
	allchan := make(chan interface{}, 3)

	allchan <- 10
	allchan <- "哈哈哈"
	allchan <- Cat{"小花", 18}

	//我们希望获得到管道中的第三个元素,则先将前2个推出
	<-allchan
	<-allchan
	//从管道中取出的cat是什么?
	newCat := <-allchan

	fmt.Printf("newCat=%T , newCat=%v\n", newCat, newCat)

	//下面的写法是错误的!编译不通过
	// fmt.Printf("newCat.Name=%v", newcat.Name)
	//使用类型断言
	fmt.Printf("newCat.Name=%v", newCat.(Cat).Name)
}
