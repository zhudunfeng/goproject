package main

import (
	"fmt"
	"go_code/chatroom/common/message"
	process2 "go_code/chatroom/server/process"
	"go_code/chatroom/server/utils"
	"io"
	"net"
)

//先创建一个processor结构体
type Processor struct {
	Conn net.Conn
}

func (this *Processor) process2() (err error) {
	//循环的客户端发送的信息
	for {
		//这里我们将读取数据包，直接封装成一个函数readPkg(),返回Message,Err
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Printf("客户端%v已断开\n", this.Conn.RemoteAddr())
				return err
			} else {
				fmt.Println("readPkg err=", err)
				//停掉当前协程
				fmt.Printf("客户端%v已断开\n", this.Conn.RemoteAddr())
				return err
			}
		}
		// fmt.Println("mes = ", mes)
		//消息总控制器
		err = this.serverProcessMes(&mes)
		if err != nil {
			return err
		}
	}
}

//编写一个ServerProcessMes 函数
//功能：根据客户端发送信息钟类不同，决定调用哪个函数来处理
func (this *Processor) serverProcessMes(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		//处理登录逻辑
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		//处理注册
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessRegister(mes)

	default:
		fmt.Println("消息类型不存在，无法处理......")
	}
	return
}
