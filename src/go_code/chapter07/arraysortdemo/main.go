package main

import "fmt"

func main() {
	numArray := [...]int{1, 4, 5, 2, 3}
	var arr1 [3]int //基本数据类型
	arr1 = [3]int{1, 2, 3}
	fmt.Printf("blblSort1:arr1[%v]=%v\n", 0, arr1[0])
	blblSort1(arr1[:]) //这里转换成切片进行传值
	fmt.Printf("blblSort1:arr1[%v]=%v\n", 0, arr1[0])

	var arr2 []int //引用数据类型
	arr2 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	fmt.Printf("blblSort1:arr2[%v]=%v\n", 0, arr2[0])
	blblSort1(arr2)
	// fmt.Println(length)
	fmt.Printf("blblSort1:arr2[%v]=%v\n", 0, arr2[0])

	// blblSort(&numArray)
	quickSort(&numArray)
	fmt.Println(numArray)
}

func blblSort1(numArray []int) int {
	numArray[0] = 20
	return len(numArray)
}

//冒泡排序
func blblSort(numArray *[5]int) {

	//控制轮数
	for i := 1; i < len(*numArray); i++ {
		//控制交换次数
		for j := 0; j < len(*numArray)-i; j++ {
			if (*numArray)[j] > (*numArray)[j+1] {
				tmp := (*numArray)[j]
				(*numArray)[j] = (*numArray)[j+1]
				(*numArray)[j+1] = tmp
			}
		}
	}

}

func quickSort(numArray *[5]int) {

	for i := 0; i < len(*numArray); i++ {
		//假定最小值
		var min int = (*numArray)[i]
		//假定最小值索引
		var index int = i
		for j := i + 1; j < len(*numArray); j++ {
			if (*numArray)[j] < min {
				min = (*numArray)[j]
				index = j
			}
		}

		if index != i {
			tmp := (*numArray)[i]
			(*numArray)[i] = (*numArray)[index]
			(*numArray)[index] = tmp
		}
	}

}
