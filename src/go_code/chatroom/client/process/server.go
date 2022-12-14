package process

import (
	"encoding/json"
	"fmt"
	"go_code/chatroom/common/message"
	"go_code/chatroom/server/utils"
	"net"
	"os"
)

//显示登录成功后的界面...
func ShowMenu() {
	fmt.Println("-------恭喜xxx登录成功--------")
	fmt.Println("-------1. 显示在线用户列表--------")
	fmt.Println("-------2. 发送消息--------")
	fmt.Println("-------3. 信息列表--------")
	fmt.Println("-------4. 退出系统--------")
	fmt.Println("请选择(1-4):")
	var key int
	var content string

	//因为，我们总会使用到SmsProcess实例，因此我们将其定义在switch外部
	smsProcess := &SmsProcess{}
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		// fmt.Println("显示在线用户列表~")
		outputOnlineUsers()
	case 2:
		// fmt.Println("发送消息")
		SelectSendType(content, smsProcess)
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("你选择退出了系统...")
		//删除server在线列表中的当前用户
		up := &UserProcess{}
		up.Logout(CurUser.UserId, CurUser.Conn)
		// time.Sleep(time.Second * 2)
		os.Exit(0)
	default:
		fmt.Println("你输入的选项不正确")
	}
}

func SelectSendType(content string, smsProcess *SmsProcess) {
	var key int
	fmt.Println("请选择发送类型(1群发,2点对点):")
	fmt.Scanf("%d\n", &key)
aa:
	for {
		switch key {
		case 1:
			fmt.Println("你想对大家说什么:)")
			fmt.Scanf("%s\n", &content)
			smsProcess.SendGroupMes(content)
			break aa
		case 2:
			var userId int
			fmt.Println("请输入目标用户的id:")
			fmt.Scanf("%d\n", &userId)
			fmt.Printf("你想对XXX说什么:)\n")
			fmt.Scanf("%s\n", &content)
			targetUser := &message.User{
				UserId: userId,
			}
			smsProcess.SendPointToPointMes(content, targetUser)
			break aa
		}
	}
}

//和服务端保持通讯
func serverProcessMes(conn net.Conn) {
	//创建一个transfer实例，不停的读取服务端发送的消息
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		fmt.Println("客户端正在等待读取服务器发送的消息")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg error: ", err)
			return
		}
		//如果读取到消息，又是下一步处理逻辑
		// fmt.Println("mes = %v\n", mes)
		switch mes.Type {
		case message.NotifyUserStatusMesType: //有人上线了
			//1.取出NotifyUserStatusMes
			var notifyUserStatusMes message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)

			//2.把这个用户的消息，状态保存到客户map[int]User中
			updateUserStatus(&notifyUserStatusMes)
			//处理
		case message.SmsMesType: //有人群发消息
			outputGroupMes(&mes)
		case message.SmsPointToPointMesType: //点对点转发
			outputPointToPointMes(&mes)
		case message.LogoutMesType: //有人登出
			outputLogoutUsers(&mes)
		default:
			fmt.Println("服务器返回了未知的消息类型")
		}
	}
}
