package main

import (
	"fmt"
	"strconv"
	"time"
)

func test() {
	for i := 1; i <= 10; i++ {
		fmt.Println("test() hello world" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main() {
	go test() //开启一个协程
	for i := 1; i <= 10; i++ {
		fmt.Println("main() hello world" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
