package main

import "fmt"

func main() {
	//打印空心三角
	/*
			*      星星：1  2*层数-1 空格：总层数-当前层数
		 ***     星星：3  2*层数-1 空格：总层数-当前层数
		*****    星星：5  2*层数-1 空格：总层数-当前层数

	*/
	var totalLevel int = 10
	//控制层数
	for i := 1; i <= totalLevel; i++ {

		//控制打印的空格数
		for k := 1; k <= totalLevel-i; k++ {
			fmt.Printf(" ")
		}

		//控制打印的星数
		for j := 1; j <= 2*i-1; j++ {
			if j == 1 || j == 2*i-1 || i == totalLevel {
				fmt.Printf("*")
			} else {
				fmt.Print(" ")
			}

		}
		//控制换行
		fmt.Println()
	}
}
