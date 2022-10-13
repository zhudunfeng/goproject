package process2

import (
	"encoding/json"
	"fmt"
	"go_code/chatroom/common/message"
	"go_code/chatroom/server/model"
	"go_code/chatroom/server/utils"
	"net"
)

type UserProcess struct {
	//字段
	Conn net.Conn

	//
}

//编写一个serverProcessLogin函数，专门处理登录请求
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	//核心代码
	//1. 先从mes 中取出 mes.Data,并直接反序列化成LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal failed:", err)
		return
	}

	//先声明一个 resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	//再声明一个LoginResMes
	var loginResMes message.LoginResMes

	//我们需要到redis数据库去完成验证
	//1. 使用model.MyUserDao到redis去验证
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {

		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误...."
		}
	} else {
		//合法
		loginResMes.Code = 200
		fmt.Println(*user, "登录成功")
	}
	// //如果用户id=100,密码=123456，认为合法，否则不合法
	// if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
	// 	//合法
	// 	loginResMes.Code = 200
	// } else {
	// 	//不合法
	// 	loginResMes.Code = 500 //500状态码，表示该用户不存在
	// 	loginResMes.Error = "该用户不存在，请注册再使用..."
	// }

	//3 将LoginResMes 序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal failed:", err)
		return
	}

	//4. 将data赋值给resMes
	resMes.Data = string(data)

	//5. 对resMes进行序列化，准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal failed:", err)
		return
	}

	//6.发送data，我们将其封装到writePkg函数
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}
