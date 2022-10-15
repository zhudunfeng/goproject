package process

import (
	"fmt"
	"go_code/chatroom/common/message"
)

//客户端要维护的map
var onlineUsers map[int]*message.User = make(map[int]*message.User, 1024)

//在客户端显示当前在线的用户
func outputOnlineUsers() {
	//遍历一把 onlineUsers
	fmt.Println("当前在线用户列表")
	for id, _ := range onlineUsers {
		fmt.Println("用户id:\t", id)
	}
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