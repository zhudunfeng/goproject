package main

import "fmt"

func main() {
	//使用for-range遍历map
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "上海"
	cities["no3"] = "广州"

	for k, v := range cities {
		fmt.Printf("k=%v v=%v\n", k, v)
	}

	//使用for-range遍历一个结构比较复杂的map
	studentMap := make(map[string]map[string]string)
	studentMap["stu01"] = make(map[string]string, 3)
	studentMap["stu01"]["name"] = "张三"
	studentMap["stu01"]["sex"] = "男"
	studentMap["stu01"]["address"] = "北京"

	studentMap["stu02"] = make(map[string]string, 3)
	studentMap["stu02"]["name"] = "Tom"
	studentMap["stu02"]["sex"] = "男"
	studentMap["stu02"]["address"] = "香港"

	for k1, v1 := range studentMap {
		fmt.Println("k1=", k1)
		for k2, v2 := range v1 {
			fmt.Printf("\tk2=%v v2=%v\n", k2, v2)
		}
		fmt.Println()
	}
}
