package message

const (
	LoginMesType            = "loginMes"
	LoginResMesType         = "loginResMes"
	RegisterMesType         = "registerMes"
	RegisterResMesType      = "registerResMes"
	NotifyUserStatusMesType = "NotifyUserStatus"
)

//这里我们定义几个用户状态的常量
const (
	UserOnline = iota
	UserOffline
	UserBusyStatus
)

type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"` //消息数据
}

//定义两个消息	后面需要再继续增加
//登录请求
type LoginMes struct {
	UserId   int    `json:"userId"`   //用户id
	UserPwd  string `json:"userPwd"`  //用户密码
	UserName string `json:"userName"` //用户名
}

//服务返回
type LoginResMes struct {
	Code    int    `json:"code"`    //返回状态码 500 表示该用户未注册 200表示成功
	UsersId []int  `json:"usersId"` //增加字段，保存用户id的切片
	Error   string `json:"error"`   //返回错误信息
}

type RegisterMes struct {
	//注册相关
	User User `json:"user"` //类型就是User结构体
}

type RegisterResMes struct {
	Code  int    `json:"code"`  //返回状态码 400 表示该用户已经占有  200表示注册成功
	Error string `json:"error"` //返回错误信息
}

//为了配合服务器推送用户状态变化的消息
type NotifyUserStatusMes struct {
	UserId int `json:"userId"` //用户id
	Status int `json:"status"` //用户的状态
}
