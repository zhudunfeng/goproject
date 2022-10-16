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

//写方法转发消息
func (this *SmsProcess) SendGroupMes(mes *message.Message) {
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
