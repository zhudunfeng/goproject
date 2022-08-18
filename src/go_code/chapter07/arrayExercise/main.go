package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 1)创建一个byte类型的26个元素的数组，分别放置'A'-Z‘。
	// 使用for循环访问所有元素并打印出来。提示:字符数据运算'A'+1 > 'B'
	var mychars [26]byte
	for i := 0; i < len(mychars); i++ {
		mychars[i] = 'A' + byte(i)
	}
	//打印
	for i := 0; i < len(mychars); i++ {
		fmt.Printf("%c ", mychars[i])
	}

	fmt.Println()

	//请求出一个数组的最大值，并得到对应的下标

	//思路
	//1.声明一个数组
	//2.假定第一个元素就是最大值，下标是0
	//3.然后从第二个元素开始循环比较，如果发现更大，则交换
	intArr := [...]int{1, -1, 9, 90, 11}
	maxValue := intArr[0]
	maxValueIndex := 0
	for i := 0; i < len(intArr); i++ {
		if maxValue < intArr[i] {
			maxValue = intArr[i]
			maxValueIndex = i
		}
	}
	fmt.Printf("maxValue=%v maxValueIndex=%v\n\n", maxValue, maxValueIndex)

	//请求出一个数组的和和平均值。for-range
	//思路
	//1.声明一个数组
	//2.求出sum
	//3.求出平均值
	intArr2 := []int{1, -1, 9, 90, 12}
	sum := 0
	for _, val := range intArr2 {
		sum += val
	}
	//如何让平均值保留到小数
	fmt.Printf("sum=%v 平均值=%v\n\n", sum, float64(sum)/float64(len(intArr2)))

	//要求：随机生成五个数，并将其反转打印
	//思路
	//1.随机生成五个数，rand.Intn()函数，注意:需要种子
	//2.当我们得到随机数后，就放在一个数组 int数组
	//3.反转打印，交换的次数是 len/2,倒数第一个和最后一个交换
	var intArr3 [5]int

	//为每次生成随机数不一样，我们需要一个seed
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(intArr3); i++ {
		intArr3[i] = rand.Intn(100)
	}

	//反转打印，交换次数是len/2
	fmt.Println("交换前=", intArr3)
	temp := 0
	for i := 0; i < len(intArr3)/2; i++ {
		temp = intArr3[len(intArr3)-1-i]
		intArr3[len(intArr3)-1-i] = intArr3[i]
		intArr3[i] = temp
	}
	fmt.Println("交换后=", intArr3)

}
