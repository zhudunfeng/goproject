package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	Num int
}

type B struct {
	Num int
}

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

func main() {
	var a A
	var b B
	//可以转换，但是有要求，就是结构体的字段要完全一样（包括：名字、个数和类型）
	a = A(b)
	fmt.Println(a, b)

	//1.创建一个Monster变量
	monster := Monster{"牛魔王", 500, "芭蕉扇~"}

	//2.将monster变量序列化为json格式字串
	//json.Marshal函数中使用反射，这个讲解反射时，我会详细介绍
	jsonStr, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("json处理错误", err)
	}
	fmt.Println("jsonStr", string(jsonStr))
}
