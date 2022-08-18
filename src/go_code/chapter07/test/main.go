package main

import (
	"fmt"
)

func main() {
	fmt.Println("哈哈")
	array := [5]int{1, 2, 5, 4, 3}
	arraySort(&array)
	fmt.Println(array)
}

func arraySort(arr *[5]int) {

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				tmp := (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = tmp
			}
		}
	}

}
