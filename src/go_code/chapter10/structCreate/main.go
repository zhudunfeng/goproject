package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	//第一种方式
	var p1 Person
	p1.Name = "Tom"
	p1.Age = 18
	fmt.Println(p1)

	//第二种方式【推荐】
	// p2 := Person{Name: "Mary", Age: 18}
	p2 := Person{"Mary", 18}
	// p2.Name = "Tom1"
	// p2.Age = 19
	fmt.Println(p2)
	//第三种方式
	var p3 *Person = new(Person)
	(*p3).Name = "Jack"
	p3.Name = "Jhon"

	(*p3).Age = 20
	p3.Age = 21
	fmt.Println(*p3)

	//第四种方式
	// var p4 *Person = &Person{"Mary",60}
	var p4 *Person = &Person{}
	(*p4).Name = "HAHA"
	p4.Name = "LiLy" //编译器底层进行了取值优化
	(*p4).Age = 24
	p4.Age = 22
	fmt.Println(*p4)
}
