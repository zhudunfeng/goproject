package main

import "fmt"

func main() {
	//定义数组
	arrays := [...]string{"宋江", "吴用", "卢俊义"}

	//普通遍历
	// for i := 0; i < len(arrays); i++ {
	// 	fmt.Printf("arrays[%d]=%s\n", i, arrays[i])
	// }

	//for-range遍历
	for index, value := range arrays {
		fmt.Printf("index=%d,value=%v\n", index, value)
	}

}
