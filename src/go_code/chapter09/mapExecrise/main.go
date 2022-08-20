package main

import "fmt"

func main() {
	//案例
	/*
		课堂练习:演示一个key-value的 value是map的案例
		比如:我们要存放3个学生信息。每个学生有name和sex信息
		思路:map[string]map[string]string
	*/
	studentMap := make(map[string]map[string]string)

	studentMap["stu01"] = make(map[string]string, 3)
	studentMap["stu01"]["name"] = "张三"
	studentMap["stu01"]["sex"] = "男"
	studentMap["stu01"]["address"] = "北京"

	studentMap["stu02"] = make(map[string]string, 3)
	studentMap["stu02"]["name"] = "Tom"
	studentMap["stu02"]["sex"] = "男"
	studentMap["stu02"]["address"] = "香港"

	studentMap["stu03"] = make(map[string]string, 3)
	studentMap["stu03"]["name"] = "Mary"
	studentMap["stu03"]["sex"] = "女"
	studentMap["stu03"]["address"] = "上海"

	fmt.Println(studentMap)
	fmt.Println(studentMap["stu02"])

}
