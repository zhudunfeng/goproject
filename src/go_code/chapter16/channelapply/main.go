package main

import (
	"fmt"
	"time"
)

//write data
func writeData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		intChan <- i
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("write data: %v\n", i)
	}
	close(intChan)
}

//read data
func readData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("readData读取到数据: %v\n", v)
	}

	//readData读取完数据后,即任务完成
	exitChan <- true
	close(exitChan)
}

func main() {
	//创建两个管道
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)
	go writeData(intChan)
	go readData(intChan, exitChan)
	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
}
