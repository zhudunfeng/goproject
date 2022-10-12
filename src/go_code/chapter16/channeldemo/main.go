package main

import "fmt"

func main() {
	//演示一下管道的使用
	//1.创建一个可以存放3个int类型的管道
	var intChan chan int
	intChan = make(chan int, 3)

	//2.看看intchan是什么
	fmt.Printf("intChan的值%v intChan本身的地址%v\n", intChan, &intChan)

	//3.向管道写入数据
	intChan <- 10
	num := 111
	intChan <- num
	// intChan <- 50
	//注意点，当我们给管写入数据时,不能超过其容量
	//fatal error: all goroutines are asleep - deadlock!
	// intChan <- 66

	//4.看看管道的长度和cap(容量)
	fmt.Printf("channel len=%v cap=%v\n", len(intChan), cap(intChan))

	//5．从管道中读取数据
	var num2 int
	num2 = <-intChan
	fmt.Printf("num2=%v\n", num2)
	fmt.Printf("channel len=%v cap=%v\n", len(intChan), cap(intChan))

	//6．在没有使用协程的情况下，如果我们的管道数据已经全部取出，再取就会报告 deadlock
	num3 := <-intChan
	//访问超过管道容量会报错
	//fatal error: all goroutines are asleep - deadlock!
	// num4 := <-intChan
	fmt.Printf("num3=%v\n", num3)
	// fmt.Printf("num3=%v num4=%v\n", num3, num4)

}
