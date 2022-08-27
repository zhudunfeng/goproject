package main

import "fmt"

/*
1)编写一个Student结构体，包含name、gender、age、id、score字段，分别为string、string、int、int、 float64类型。
2）结构体中声明一个say方法，返回string类型，方法返回信息中包含所有字段值。
3)在main方法中，创建 Student结构体实例(变量)，并访问say方法，并将调用结果打印输出。

*/
type Student struct {
	name   string
	gender string
	age    int
	id     int
	score  float64
}

func (s *Student) say() string {
	infoStr := fmt.Sprintf("student的信息是 name=[%v] gender=[%v] age=[%v] id=[%v] score=[%v]",
		s.name, s.gender, s.age, s.id, s.score)
	return infoStr
}

func main() {
	stu := Student{
		name:   "tom",
		gender: "男",
		age:    18,
		id:     1000,
		score:  99.99,
	}
	infoStr := stu.say()
	fmt.Println(infoStr)

	//测试长方体
	// var box Box
	// fmt.Println("请输入len:")
	// fmt.Scanln(&(box.len))
	// fmt.Println("请输入width:")
	// fmt.Scanln(&(box.width))
	// fmt.Println("请输入height:")
	// fmt.Scanln(&(box.height))
	// volumn := box.getVolumn()
	// fmt.Printf("volumn=%.2f\n", volumn)

	//测试游客案例
	var visitor Visitor
	for {
		fmt.Println("请输入你的名字")
		fmt.Scanln(&visitor.Name)
		if visitor.Name == "n" {
			fmt.Println("退出程序...")
			break
		}
		fmt.Println("请输入你的年龄")
		fmt.Scanln(&visitor.Age)
		visitor.showPrice()
	}
}
