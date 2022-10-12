package main

import (
	"fmt"
	"go_code/chatroom/common/message"
	"net"
)

//编写一个ServerProcessMes 函数
//功能：根据客户端发送信息钟类不同，决定调用哪个函数来处理
func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		//处理登录逻辑
		err = serverProcessLogin(conn, mes)
	case message.RegisterMesType:
		//处理注册

	default:
		fmt.Println("消息类型不存在，无法处理......")
	}
	return
}
