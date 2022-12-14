package process2

import (
	"encoding/json"
	"fmt"
	"go_code/chatroom/common/message"
	"go_code/chatroom/server/utils"
	"net"
)

type SmsProcess struct {
}

//点对点消息
func (this *SmsProcess) SendPointToPointMes(mes *message.Message) {
	var smsPointToPointMes message.SmsPointToPointMes
	err := json.Unmarshal([]byte(mes.Data), &smsPointToPointMes)
	if err != nil {
		fmt.Println("json.Unmarshal error:", err)
		return
	}

	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal error:", err)
	}

	for id, up := range userMgr.onlineUsers {
		//这里，转发给目标用户
		if id == smsPointToPointMes.TargetUser.UserId {
			this.sendMesEachOnlineUser(data, up.Conn)
			break
		}
	}
}

//写方法转发消息
func (this *SmsProcess) SendGroupMes(mes *message.Message) {
	//遍历服务器端的onlineUsers map[int]*UserProcess
	//将消息转发取出

	//取出mes的内容 SmsMes
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal error:", err)
		return
	}

	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal error:", err)
	}

	for id, up := range userMgr.onlineUsers {
		//这里，还需要过滤自己，即不要再发给自己
		if id == smsMes.UserId {
			continue
		}
		this.sendMesEachOnlineUser(data, up.Conn)
	}

}

func (this *SmsProcess) sendMesEachOnlineUser(data []byte, conn net.Conn) {
	//创建一个Transfer实例，发送data
	tf := &utils.Transfer{
		Conn: conn,
	}

	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("转发消息失败 error: ", err)
	}
}
