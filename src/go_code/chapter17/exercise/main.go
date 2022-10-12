package main

import (
	"fmt"
	"reflect"
)

// func main() {ddd
// 	var str string = "tom" //ok
// 	fs := reflect.ValueOf(str) //ok fs-> string类型
// 	fs.SetString("Jack") //err
// 	fmt.Println(fs.String()) //jack
// }

func main() {
	var str string = "tom"      //ok
	fs := reflect.ValueOf(&str) //ok fs-> string类型
	fs.Elem().SetString("Jack") //err
	fmt.Printf("%s\n", str) //jack
	fmt.Println(fs.Elem().String()) //jack
}
