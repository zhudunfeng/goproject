package main

import "fmt"

//Stu结构体
type Stu struct {
	Name    string
	Age     int
	Address string
}

func main() {
	//map的value也经常使用strult类型，
	//更适合管理复杂的数据(比前面value是一个map更好)，
	//比如value为 student结构体【案例演示，因为还没有学结构体，体验一下即可】
	//1.map的key 为学生的学号，是唯一的
	//2.map 的value为结构体,包含学生的名字,年龄,地址
	students := make(map[string]Stu, 10)
	//创建2个学生
	stu1 := Stu{"tom", 18, "北京"}
	stu2 := Stu{"mary", 18, "上海"}
	students["no1"] = stu1
	students["no2"] = stu2

	//遍历各个学生信息
	for k, v := range students {
		fmt.Printf("学生的编号是%v\n", k)
		fmt.Printf("学生的姓名是%v\n", v.Name)
		fmt.Printf("学生的年龄是%v\n", v.Age)
		fmt.Printf("学生的地址是%v\n", v.Address)
		fmt.Println()
	}
}
