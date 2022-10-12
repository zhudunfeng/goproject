package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_code/chatroom/common/message"
	"net"
)

func login(userId int, userPwd string) (err error) {
	//TODO 下一步就要开始定义协议
	// fmt.Printf("userId=%d, userPwd=%s\n", userId, userPwd)
	// return nil

	//1.链接到服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("net.Dial error: ", err)
		return
	}

	//延迟关闭链接
	defer conn.Close()

	//2.准备通过conn发送消息给服务
	var mes message.Message
	mes.Type = message.LoginMesType
	//3.创建一个LoginMes结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	//4.将loginMes 序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal error:", err)
		return
	}

	//5.把data赋给 mes.Data字段
	mes.Data = string(data)

	//6.将mes进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal error:", err)
		return
	}

	//7.到这个时候 data就是我们要发送的消息
	//7.1 先把data的长度发送给服务器
	//先获取到data的长度->转成一个表示长度的byte切片
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)

	//发送长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes[:]) fail:", err)
		return
	}

	// fmt.Printf("客户端,发送消息的长度=%d,数据内容=%s\n", len(data), string(data))
	//真正的发送数据
	//发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) failed: ", err)
		//此时默认返回错误
		return
	}

	//TODO 这里还需要处理服务器端返回的消息

	//延时关闭客户端
	// time.Sleep(20 * time.Second)
	// fmt.Println("休眠了20s")

	mes, err = readPkg(conn)
	if err != nil {
		fmt.Println("readPkg(conn) error: ", err)
		return
	}

	//将mes的Data部分反序列化成LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)

	if loginResMes.Code == 200 {
		fmt.Println("登录成功")
	} else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
	}
	return 
}
