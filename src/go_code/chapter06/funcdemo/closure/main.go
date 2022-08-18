package main

import (
	"fmt"
	"strings"
)

//闭包的最佳实践
func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func getSumAndSub(n1 int, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}

func getSumAndSub1(n1 int, n2 int) (sum int ,sub int){
	sum = n1 + n2
	sub = n1 - n2
	return sum, sub
} 

func main() {
	f := makeSuffix(".jpg")
	fmt.Println("文件名处理后", f("winter"))
	fmt.Println("文件名处理后", f("hello.jpg"))

	n1 := 10
	n2 := 5
	sum, sub := getSumAndSub(n1, n2)
	fmt.Println("sum=", sum)
	fmt.Println("sub=", sub)
	sum1, sub1 := getSumAndSub1(n1, n2)
	fmt.Println("sum1=",sum1)
	fmt.Println("sub1=",sub1)
}
