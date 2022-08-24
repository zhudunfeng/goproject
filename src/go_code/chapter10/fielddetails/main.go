package main

import "fmt"

//如果结构体的字段类型是:指针，slice，和map的零值都是nil ，即还没有分配空间
//如果需要使用这样的字段,需要先make,才能使用.
type Person struct {
	Name   string
	Age    int
	Scores [5]int
	ptr    *int              //指针
	slice  []int             //切片
	map1   map[string]string //map
}

func main() {
	//定义结构体遍历
	var p1 Person
	fmt.Println(p1)
	if p1.ptr == nil {
		fmt.Println("ok1")
	}
	if p1.slice == nil {
		fmt.Println("ok2")
	}
	if p1.map1 == nil {
		fmt.Println("ok3")
	}

	//使用slice,再次说明，一定要make
	p1.slice = make([]int, 10)
	p1.slice[0] = 100

	//使用map一定要先make
	p1.map1 = make(map[string]string)
	p1.map1["key1"] = "mary~"
	fmt.Println(p1)
}
