package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//将../abc.txt 文件内容导入到  ../kkk.txt
	//1. 首先将  ../abc.txt 内容读取到内存
	//2. 将读取到的内容写入 ../kkk.txt
	filepath1 := "../abc.txt"
	filepath2 := "../kkk.txt"

	data, err := ioutil.ReadFile(filepath1)
	if err != nil {
		//说明读取文件有错误
		fmt.Println("open file error: ", err)
		return
	}

	err = ioutil.WriteFile(filepath2, data, 777)
	if err != nil {
		fmt.Println("write file error: ", err)
	}
}
