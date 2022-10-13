package model

import (
	"encoding/json"
	"fmt"

	"github.com/gomodule/redigo/redis"
)

//定义一个UserDao结构体体
//完成对User结构体的各种操作.
type UserDao struct {
	pool *redis.Pool
}

//我们在服务器启动后，就初始化一个userDao实例，
//把它做成全局的变量，在需要和redis操作时，就直接使用即可
var (
	MyUserDao *UserDao
)

//使用工厂模式，创建一个UserDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	return &UserDao{pool: pool}
}

//思考一下在UserDao 应该提供哪些方法给我们
//1.根据用户id返回一个User实例+err
func (this *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {
	//通过给定的id 去redis查询整个用户
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		//错误
		if err == redis.ErrNil { //表示在users哈希中，没有找打对应id
			err = ERROR_USER_NOTEXISTS
		}
		return
	}

	user = &User{}
	//这里我们需要把res 反序列化成User实例
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal error:", err)
		return
	}
	return
}

//完成登录的校验 Login
//1. Login 完成对用户的验证
//2. 如果用户的id和pwd都正确，则返回一个user实例
//3. 如果用户的id或pwd有错误，则返回对应的错误信息
func (this *UserDao) Login(userId int, userPwd string) (user *User, err error) {
	//先从UserDao的连接词中取出一根连接
	conn := this.pool.Get()
	defer conn.Close()
	user, err = this.getUserById(conn, userId)
	if err != nil {
		return
	}

	//这是证明这个用户是获取到的
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}