package main

import "fmt"

func main() {
	//演示切片的使用
	/*
		要求:使用一个map来记录monster的信息 name 和 age，
		也就是说一个monster对应一个map,并且妖怪的个数可以动态的增加=>map切片
	*/
	// 1.声明一个map切片
	//var monsters []map[string]string
	monsters := make([]map[string]string, 2)
	// 2.增加妖怪信息
	if monsters[0] == nil {
		monsters[0] = make(map[string]string, 2)
		monsters[0]["name"] = "牛魔王"
		monsters[0]["age"] = "500"
	}

	if monsters[1] == nil {
		monsters[1] = make(map[string]string, 2)
		monsters[1]["name"] = "孙悟空"
		monsters[1]["age"] = "500"
	}

	//panic: runtime error: index out of range [2] with length 2
	// if monsters[2] == nil {
	// 	monsters[2] = make(map[string]string, 2)
	// 	monsters[2]["name"] = "嫦娥"
	// 	monsters[2]["age"] = "500"
	// }

	//1.先定义一个monster信息
	newMonster := map[string]string{
		"name": "嫦娥",
		"age":  "600",
	}
	monsters = append(monsters, newMonster)

	fmt.Println(monsters)
	for index, val := range monsters {
		fmt.Println(index, val)
	}
}
