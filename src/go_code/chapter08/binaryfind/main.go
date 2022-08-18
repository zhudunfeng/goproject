package main

import "fmt"

/*
二分查找的思路:比如我们要查找的数是findval
1. arr是一个有序数组，并且是从小到大排序
2．先找到中间的下标 middle = (leftIndex + rightIndex)/ 2，然后让中间下标的值和findval进
	2.1 如果 arr[middle] > findval ,就应该向 leftIndex ---- (middle - 1)
	2.2如果 arr[middle] < findval ,就应该向 middel+1---- rightIndex
	2.3如果arr[middle] == findval ，就找到
	2.4上面的2.1 2.2 2.3 的逻辑会递归执行
3．想一下，怎么样的情况下，就说明找不到[分析出退出递归的条件!!]
	if leftIndex > rightIndex {
		/找不到..
		return ..
	}
*/
func BinaryFind(arr *[6]int, leftIndex int, rightIndex int, findval int) {
	//如果 leftIndex > rightIndex 找不到了,退出递归
	if leftIndex > rightIndex {
		fmt.Printf("找不到%v\n", findval)
		return
	}

	//先找到中间的下标
	middle := (leftIndex + rightIndex) / 2

	if (*arr)[middle] > findval {
		// 说明查找的数在 leftIndex ----middle-1
		BinaryFind(arr, leftIndex, middle-1, findval)
	} else if (*arr)[middle] < findval {
		//说明查找的数字在 middle+1 ----rightIndex
		BinaryFind(arr, middle+1, rightIndex, findval)
	} else {
		//找到了
		fmt.Printf("找到了,下标为%v\n", middle)
	}
}

func main() {
	arr := [6]int{1, 8, 10, 89, 1000, 1234}
	//测试一把
	BinaryFind(&arr, 0, len(arr)-1, 89)
}
