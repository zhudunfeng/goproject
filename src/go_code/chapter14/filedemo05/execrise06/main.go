package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//自己编写一个函数，接收两个文件路径 srcFileName dstFileName
func Copy(dstFileName string, srcFileName string) (written int64, err error) {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Println("open src file err: ", err)
		return
	}
	defer srcFile.Close()
	//通过srcfile ,获取到 Reader
	reader := bufio.NewReader(srcFile)

	//打开dstFileName
	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 777)
	if err != nil {
		fmt.Println("open dst file err: ", err)
	}
	defer dstFile.Close()
	//通过dstFile, 获取到 Writer
	writer := bufio.NewWriter(dstFile)
	return io.Copy(writer, reader)
}

func main() {
	//将C:\\Users\\21015\\Downloads\\test333(3)_1.png 文件拷贝到 e:/abc.jpg
	//调用../abc.png 完成文件拷贝
	srcFileName := "C:\\Users\\21015\\Downloads\\test333(3)_1.png"
	dstFileName := "../abc.png"
	_, err := Copy(dstFileName, srcFileName)
	if err == nil {
		fmt.Println("拷贝完成")
	} else {
		fmt.Println("拷贝错误 error: ", err)
	}
}
