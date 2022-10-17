package process2

import "fmt"

//因为UserMgr 实例在服务器端有且只有一个
//因为在很多的地方，都会使用到，因此，我们
//将其定义为全局变量
var (
	userMgr *UserMgr
)

type UserMgr struct {
	//在线用户
	onlineUsers map[int]*UserProcess
	//离线用户id列表
	offlineUsers []int
}

//完成对userMgr的初始化工作
func init() {
	userMgr = &UserMgr{
		onlineUsers:  make(map[int]*UserProcess, 1024),
		offlineUsers: make([]int, 0),
	}
}

//完成对onlineUsers添加
func (this *UserMgr) AddOnlineUsers(up *UserProcess) {
	this.onlineUsers[up.UserId] = up
}

//删除
func (this *UserMgr) DelOnlineUsers(userId int) {
	delete(this.onlineUsers, userId)
}

//返回当前所有在线的用户
func (this *UserMgr) GetAllOnlineUsers() map[int]*UserProcess {
	return this.onlineUsers
}

//根据当前id返回对应的值
func (this *UserMgr) GetOnlineUserById(userId int) (up *UserProcess, err error) {
	//从map取出一个值，带检测方式
	up, ok := this.onlineUsers[userId]
	if !ok { //说明，你要查找的这个用户，当前不在线
		err = fmt.Errorf("用户%d不存在\n", userId)
		return
	}
	return
}
