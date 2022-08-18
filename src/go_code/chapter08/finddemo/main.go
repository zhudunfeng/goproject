package main

import "fmt"

func main() {
	//有一个数列:白眉鹰王、金毛狮王、紫衫龙王、青翼蝠王
	//猜数游戏:从键盘中任意输入一个名称，判断数列中是否包含此名称【顺序查找】
	//思路
	//1定义一个数组，白眉鹰王、金毛狮王、紫衫龙王、青翼蝠王字符串数组
	//2.从控制台接收一个名字，依次比较,如果发现有,提示
	names := [5]string{"白眉鹰王", "金毛狮王", "紫衫龙王", "青翼蝠王"}
	var heroname string = ""
	fmt.Println("请输入要查找的人名...")
	fmt.Scanln(&heroname)

	//顺序查找:第一种方式
	for i := 0; i < len(names); i++ {
		if heroname == names[i] {
			fmt.Printf("找到了%v 下标%v\n", names[i], i)
			break
		} else if i == len(names)-1 {
			fmt.Printf("没找到了%v\n", heroname)
		}
	}

	//顺序查找:第2种方式（推荐。。。）
	index := -1
	for i := 0; i < len(names); i++ {
		if heroname == names[i] {
			index = i
		}
	}
	if index != -1 {
		fmt.Printf("找到了%v 下标%v\n", names[index], index)
	} else {
		fmt.Printf("没找到了%v\n", heroname)
	}
}
