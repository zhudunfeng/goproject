package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
)

// 实现对Hero结构体切片的排序: sort.Sort(data Interface)
//1.声明Hero结构体
type Hero struct {
	Name string
	Age  int
}

//2.声明一个Hero结构体切片类型
type HeroSlice []Hero

//3.实现Interface 接口
func (hs HeroSlice) Len() int { return len(hs) }

func (hs HeroSlice) Less(i, j int) bool {
	//按年龄升序
	// return hs[i].Age < hs[j].Age
	//按名字acii升序
	return hs[i].Name < hs[j].Name
}

func (hs HeroSlice) Swap(i, j int) {
	//交换
	hs[i], hs[j] = hs[j], hs[i]
}

func main() {
	//先定义一个数组/切片
	var intSlice = []int{10, -1, 5, 3, 7}
	//要求对 intSlice切片进行排序
	//1. 冒泡排序...
	//2. 也可以使用系统提供的方法
	fmt.Println(intSlice)
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	//请大家对结构体切片进行排序
	//1. 冒泡排序...
	//2. 也可以使用系统提供的方法
	//测试看看我们是否可以对结构体切片进行排序
	var hreos HeroSlice

	for i := 0; i < 10; i++ {
		hreo := Hero{
			Name: fmt.Sprintf("英雄|%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		//将 hero append到 heroes切片
		hreos = append(hreos, hreo)
	}

	fmt.Println("=============================")
	//看看排序前的顺序
	for _, v := range hreos {
		fmt.Println("Name=", v.Name, "Age=", v.Age)
	}

	fmt.Println("=============================")
	//排序，接口和切片都是引用类型
	sort.Sort(hreos)

	//看看排序后的顺序
	for _, v := range hreos {
		fmt.Println("Name=", v.Name, "Age=", v.Age)
	}
	fmt.Println("=============================")

	var students StudentSlice

	for i := 0; i < 10; i++ {
		score, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", 100*rand.Float64()), 64)
		student := Student{
			Name:  fmt.Sprintf("学生%d", rand.Intn(100)),
			Age:   rand.Intn(100),
			Score: score,
		}
		students = append(students, student)
	}
	sort.Sort(students)

	for _, v := range students {
		fmt.Println(v)
	}

}
