package main

import (
	"fmt"
	"go_code/chapter11/encapexercise/model"
)

func main() {
	var account = model.NewAccount("Adun1314", "666666", 200000)

	if account == nil {
		fmt.Println("创建失败...")
	} else {
		fmt.Println(*account)
		fmt.Println("创建成功")
	}
}
