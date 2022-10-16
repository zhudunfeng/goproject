package process

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_code/chatroom/common/message"
	"go_code/chatroom/server/utils"
	"net"
	"os"
)

type UserProcess struct {
}

func (this *UserProcess) Register(userId int, userPwd string, userName string) (err error) {

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
	mes.Type = message.RegisterMesType
	//3.创建一个LoginMes结构体
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	//4.将loginMes 序列化
	data, err := json.Marshal(registerMes)
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

	//创建一个tranfer
	tf := &utils.Transfer{
		Conn: conn,
	}

	//发送data给服务器端
	err = tf.WritePkg(data)

	if err != nil {
		fmt.Println("注册发送消息错误 error:", err)
	}

	//这里还需要处理服务器端返回的消息
	mes, err = tf.ReadPkg() //mes 就是RegisterMes
	if err != nil {
		fmt.Println("readPkg(conn) error: ", err)
		return
	}

	//将mes的Data部分反序列化成LoginResMes
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)

	if registerResMes.Code == 200 {
		fmt.Println("注册成功，你重新登录一把")
		os.Exit(0)
	} else {
		fmt.Println(registerResMes.Error)
		os.Exit(0)
	}
	return
}

func (this *UserProcess) Login(userId int, userPwd string) (err error) {
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
	tf := &utils.Transfer{
		Conn: conn,
	}
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg(conn) error: ", err)
		return
	}

	//将mes的Data部分反序列化成LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)

	if loginResMes.Code == 200 {
		// fmt.Println("登录成功")
		//初始化
		CurUser.Conn = conn
		CurUser.UserId = userId
		CurUser.UserStatus = message.UserOnline

		fmt.Println("当前在线用户列表如下:")
		for _, v := range loginResMes.UsersId {
			//如果我们要求不显示自己在线，下面我们增加一个代码
			if v == userId {
				continue
			}
			fmt.Println("用户id:\t", v)
			//完成 客户端的onlineUsers 初始化
			user := &message.User{
				UserId:     userId,
				UserStatus: message.UserOnline,
			}
			onlineUsers[v] = user
		}
		fmt.Print("\n\n")
		//可以显示当前在线用户列表，遍历loginResMes.UsersId
		//这里我们需要在客户端启动一个协程
		//该协程保持和服务器端的通讯，如果服务器有数据推送给客户端
		//则接收并显示在客户端的终端
		go serverProcessMes(conn)
		//1. 显示我们的登录成功菜单[循环]
		for {
			ShowMenu()
		}
	} else {
		fmt.Println(loginResMes.Error)
	}
	return
}
