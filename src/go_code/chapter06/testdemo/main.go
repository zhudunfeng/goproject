package main
import (
	"fmt"
	"os"
)

func main() {
	//打开文件
	file, err := os.Open("E:/test.txt")
	if err != nil {
		fmt.Println(err)
	}
	//读取文件内容
	data := make([]byte, 100)
	n, err := file.Read(data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data[:n]))
	
	file, err = os.OpenFile("E:/test.txt", os.O_RDWR|os.O_CREATE, 0666)
	//写入文件内容
	file.WriteString("hello go")
	//关闭文件
	file.Close()
}