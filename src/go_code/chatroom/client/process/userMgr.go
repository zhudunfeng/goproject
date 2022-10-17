package process

import (
	"encoding/json"
	"fmt"
	"go_code/chatroom/client/model"
	"go_code/chatroom/common/message"
)

//客户端要维护的map
var onlineUsers map[int]*message.User = make(map[int]*message.User, 1024)
var CurUser model.CurUser //我们在客户登录成功后，完成对CurUser初始化

//在客户端显示当前在线的用户
func outputOnlineUsers() {
	//遍历一把 onlineUsers
	fmt.Println("当前在线用户列表")
	for id, _ := range onlineUsers {
		fmt.Println("用户id:\t", id)
	}
}

func outputLogoutUsers(mes *message.Message) { //接收用户登出消息
	//显示即可
	//1.反序列化mes.Data
	var logoutMes message.LogoutMes
	err := json.Unmarshal([]byte(mes.Data), &logoutMes)
	if err != nil {
		fmt.Println("json.Unmarshal error:", err.Error())
		return
	}


	delete(onlineUsers,logoutMes.UserId)
	
	//打印当前在线用户
	outputOnlineUsers()
}

//编写一个方法，处理返回的NotifyUserStatusMes
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {
	//适当优化
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok {
		user = &message.User{
			UserId: notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId] = user
	//打印
	outputOnlineUsers()
}
