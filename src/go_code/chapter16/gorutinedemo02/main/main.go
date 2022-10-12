package main

import (
	"fmt"
	"sync"
	"time"
)

//思路
// 1。编写一个函数,来计算各个数的阶乘，并放入到 map中.
// 2．我们启动的协程多个,统计的将结果放入到map中
// 3. map应该做出一个全局的.

var (
	myMap = make(map[int]int, 10)
	//声明一个全局的互斥锁
	//lock是一个全局的互斥锁,
	//sync是包:synchornized同步
	//Mutex :是互斥
	lock sync.Mutex
)

func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= n
	}
	//这里我们将res放入到myMap
	//加锁
	lock.Lock()
	myMap[n] = res //fatal error: concurrent map writes
	//解锁
	lock.Unlock()
}

func main() {
	//我们这里开启多个协程完成这个任务[200个]
	for i := 1; i <= 200; i++ {
		go test(i)
	}

	// 休眠10秒种【等多久】
	time.Sleep(10 * time.Second)

	//这里我们输出结果,变量这个结果
	//加锁
	lock.Lock()
	for k, v := range myMap {
		fmt.Printf("map[%d] = %d\n", k, v)
	}
	//解锁
	lock.Unlock()

}
