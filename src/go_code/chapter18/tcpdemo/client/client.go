package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.0.8:8888")
	if err != nil {
		fmt.Println("client dial error: ", err)
		return
	}

	//功能一：客户端可以发送单行数据，然后退出
	reader := bufio.NewReader(os.Stdin) //os.Stdin 代表标准输入

	for {
		//从终端读取一行用户输入，并准备发送服务器
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("readString error: ", err)
		}

		//如果用户输入的是 exit就退出
		line = strings.Trim(line, " \r\n")
		if line == "exit" {
			fmt.Println("客户端退出......")
			break
		}
		//再将line 发送给服务器
		_, err = conn.Write([]byte(line + "\n"))
		if err != nil {
			fmt.Println("conn.Write error: ", err)
		}
	}
	// fmt.Printf("客户端发送了%d字节的数据，并退出", n)
}
