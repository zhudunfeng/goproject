package main

import (
	"fmt"
	"time"
)

func main() {
	//使用select可以解决从管道去数据的阻塞问题

	//1.定义一个管道 10个数据int
	intChan := make(chan int, 10)

	for i := 0; i < 10; i++ {
		intChan <- i
	}

	//2.定义一个管道 5个数据string
	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}

	//传统的方法在遍历管道时，如果不关闭会阻塞而导致  deadlock

	//问题，在实际开发中，可能我们不好确定什么时候关闭该管道
	//可以使用select 方式解决
label:
	for {
		select {
		case v := <-intChan:
			fmt.Printf("从intChan读取的数据%d\n", v)
			time.Sleep(100 * time.Microsecond)
		case v := <-stringChan:
			fmt.Printf("从stringChan读取的数据%s\n", v)
			time.Sleep(100 * time.Microsecond)
		default:
			fmt.Printf("都取不到,不玩了\n")
			// return
			break label
		}
	}
}
