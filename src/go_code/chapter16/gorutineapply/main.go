package main

import (
	"fmt"
	"time"
)

//向intchan放入1-8000个数
func putNum(intChan chan int) {
	for i := 1; i <= 80; i++ {
		intChan <- i
	}
	//关闭intchan
	close(intChan)
}

//从intchan取出数据,并判断是否为素数,如果是,就
//放入到primechan
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	//使用for循环
	var flag bool
	for {
		time.Sleep(10 * time.Millisecond)
		num, ok := <-intChan
		flag = true //假设是素数
		if !ok {
			//intChan管道中取不到数，退出循环
			break
		}
		//判断num是不是素数
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false //该数不是素数
				break
			}
		}

		//将这个数就放入到primechan
		if flag {
			primeChan <- num
		}
	}

	fmt.Println("有一个primeNum协程因为取不到数据,退出")
	//这里我们还不能关闭primechan
	//向exitchan写入true
	exitChan <- true
}

func main() {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 2000) //放入结果
	//标识退出的管道 此处标识4个协程的状态
	exitChan := make(chan bool, 4)

	//开启一个协程,向intchan放入1-8000个数
	go putNum(intChan)
	//开启4个协程,从intChan取出数据,并判断是否为素数
	for i := 0; i < 4; i++ {
		//放入到primechan
		go primeNum(intChan, primeChan, exitChan)
	}

	//这里我们主线程,进行处理
	//直接协程加立即执行函数
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		//当我们从exitchan取出了4个结果,就可以放心的关闭primeNum
		close(primeChan)
		close(exitChan)
	}()

	//遍历我们的prpimeNum ,把结果取出
	for {
		res, ok := <-primeChan
		if !ok {
			break
		}
		//将结果输出
		fmt.Printf("素数=%d\n", res)
	}
	fmt.Println("主线程退出")

}
