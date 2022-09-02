package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//打开一个存在的文件，在原来的内容追加内容 'ABC! ENGLISH!'代码实现:
	//1 .打开文件d:/abc.txt
	filepath := "../abc.txt"
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file err:", err)
	}
	//及时关闭file句柄
	defer file.Close()
	//准备写入5句 "hello, Gardon"
	str := "ABC, English\r\n" //\r\n 表示换行
	//写入时，使用带缓存的 *Writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	//因为writer是带缓存，因此在调用WriterString方法时，其实
	//内容是先写入到缓存的,所以需要调用Flush方法，将缓冲的数据
	//真正写入到文件中， 否则文件中会没有数据!!!
	writer.Flush()
}
