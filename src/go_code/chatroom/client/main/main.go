package main

import (
	"fmt"
	"go_code/chatroom/client/process"
	"math/rand"
	"os"
	"time"
)

//定义里那个变量，一个表示用户id，一个表示用户密码
var userId int
var userPwd string
var userName string

func init() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("客户端", rand.Int(), "启动成功......")
}

func main() {

	//接收用户的选择
	var key int
	//判断是否继续显示菜单
	// var loop bool = true

	for true {
		fmt.Println("-----------------欢迎登录多人聊天系统-----------------")
		fmt.Println("\t\t\t 1 登录聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 请选择(1-3):")

		fmt.Scanf("%d\n", &key)

		switch key {
		case 1:
			fmt.Println("登录聊天室")
			fmt.Println("请输入用户的id:")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户的密码:")
			fmt.Scanf("%s\n", &userPwd)
			//完成登录
			//1.创建一个UserProcess的实例
			up := &process.UserProcess{}
			up.Login(userId, userPwd)
			// loop = false
		case 2:
			fmt.Println("注册用户")
			fmt.Println("请输入用户的id:")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户的密码:")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("请输入用户名[nickname]:")
			fmt.Scanf("%s\n", &userName)
			up := &process.UserProcess{}
			up.Register(userId, userPwd, userName)
			// loop = false
		case 3:
			fmt.Println("退出系统")
			// loop = false
			os.Exit(0)
		default:
			fmt.Println("您输入有误，请重新输入")
		}

	}

	//根据用户输入，显示新的提示信息
	// if key == 1 {
	// 	//说明用户要登录
	// 	fmt.Println("请输入用户的id:")
	// 	fmt.Scanf("%d\n", &userId)
	// 	fmt.Println("请输入用户的密码:")
	// 	fmt.Scanf("%s\n", &userPwd)
	// 	//先把登录的函数，写到另外一个文件，比如login.go
	// 	// login(userId, userPwd)
	// 	// if err != nil {
	// 	// 	fmt.Println("登录失败")
	// 	// } else {
	// 	// 	fmt.Println("登录成功")
	// 	// }
	// } else if key == 2 {
	// 	fmt.Println("进行用户注册的逻辑.......")
	// }
}
