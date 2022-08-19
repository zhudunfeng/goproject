package main

import (
	"fmt"
	"time"
)

func main() {
	//3)获取到当前时间的方法:
	now := time.Now()
	fmt.Printf("now=%v now type=%T\n", now, now)

	//4)如何获取到其它的日期信息
	fmt.Printf("年=%v\n", now.Year())
	fmt.Printf("月=%v\n", now.Month())
	fmt.Printf("月=%v\n", int(now.Month()))
	fmt.Printf("日=%v\n", now.Day())
	fmt.Printf("时=%v\n", now.Hour())
	fmt.Printf("分=%v\n", now.Minute())
	fmt.Printf("秒=%v\n", now.Second())

	//格式化日期时间
	fmt.Printf("当前年月日%d-%d-%d %d :%d :%d \n",
		now.Year(), now.Month(), now.Day(), now.Hour(),
		now.Minute(), now.Second())

	dataStr := fmt.Sprintf("当前年月日%d-%d-%d %d :%d :%d \n",
		now.Year(), now.Month(), now.Day(), now.Hour(),
		now.Minute(), now.Second())

	fmt.Printf("dataStr=%v\n", dataStr)

	//格式化日期时间的第二种方式
	fmt.Println(now.Format("2006/01/02 15:04:05"))
	fmt.Println()
	fmt.Println(now.Format("2006-01-02"))
	fmt.Println()
	fmt.Println(now.Format("15:04:05"))
	fmt.Println()
}
