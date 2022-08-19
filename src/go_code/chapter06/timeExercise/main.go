package main

import (
	"fmt"
	"strconv"
	"time"
)

func test03() {
	str := ""
	for i := 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
}

func main() {
	//在执行testo3前;先获取到当前的unix时间戳
	start := time.Now().Unix()
	test03()
	end := time.Now().Unix()
	fmt.Printf("指定test03耗费的时间为%v秒\n", end-start)
}
