package process

import (
	"encoding/json"
	"fmt"
	"go_code/chatroom/common/message"
	"go_code/chatroom/server/utils"
)

type SmsProcess struct{}

//发送群聊的消息
func (this *SmsProcess) SendGroupMes(content string) (err error) {
	//1.创建一个Mes
	var mes message.Message
	mes.Type = message.SmsMesType

	//2.创建一个SmsMes实例
	var smsMes message.SmsMes
	smsMes.Content = content //内容
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus

	//3.序列化 smsMes
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal error:", err.Error())
		return
	}

	mes.Data = string(data)

	//4.对mes再次序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal error:", err.Error())
		return
	}
	//5.将mes发送给服务器
	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}

	//6.发送
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendGroupMes error:", err.Error())
		return
	}
	return
}

//点对点发送
func (this *SmsProcess) SendPointToPointMes(content string, targetUser *message.User) (err error) {
	//1.创建一个Mes
	var mes message.Message
	mes.Type = message.SmsPointToPointMesType

	//2.创建一个SmsPointToPointMes实例
	var smsPointToPointMes message.SmsPointToPointMes
	smsPointToPointMes.Content = content
	smsPointToPointMes.SourceUser.UserId = CurUser.UserId
	smsPointToPointMes.SourceUser.UserStatus = CurUser.UserStatus

	smsPointToPointMes.TargetUser.UserId = targetUser.UserId

	//3.序列化 smsPointToPointMes
	data, err := json.Marshal(smsPointToPointMes)
	if err != nil {
		fmt.Println("SendPointToPointMes json.Marshal error:", err.Error())
		return
	}

	mes.Data = string(data)

	//4.对mes再次序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("SendPointToPointMes json.Marshal error:", err.Error())
		return
	}

	//5.将mes发送给服务器
	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}

	//6.发送
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendPointToPointMes error:", err.Error())
		return
	}
	return
}
