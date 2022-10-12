package main

import (
	"fmt"
	_ "testing"
	"time"
)

// func TestChan(t *testing.T) {
// 	aa := make(chan interface{}, 1)
// 	aa <- "hello"
// 	close(aa)
// 	data := <-aa
// 	fmt.Println(data.(string))
// }

func Test01() {
	aa := make(chan interface{}, 1)
	aa <- "hello"
	close(aa)
	data := <-aa
	fmt.Println(data.(string))
}

func main(){
	go Test01()

	time.Sleep(time.Second*3)
}