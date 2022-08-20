package main

import "fmt"

func main() {
	//方式一
	var a map[string]string
	a = make(map[string]string, 10)
	a["no1"] = "宋江"
	a["no2"] = "吴用"
	a["no1"] = "武松"
	a["no3"] = "吴用"
	fmt.Println(a)

	//方式二
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "上海"
	fmt.Println(cities)

	//方式三
	// var heros = map[string]string{}
	// var heros map[string]string=map[string]string{}
	heros := map[string]string{
		"hero1": "宋江",
		"hero2": "卢俊义",
	}
	heros["hero3"] = "李逵"
	fmt.Println(heros)
}
