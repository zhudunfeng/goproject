package main

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

//定义全局pool
var pool *redis.Pool

//每一个源文件都可以包含l一个init函数，该函数会在main 函数执行前，
//被Go运行框架调用，也就是说init会在main函数前被调用
//当启动程序是就初始化pool
func init() {
	pool = &redis.Pool{
		MaxIdle:     8,                 //最大空闲数
		MaxActive:   0,                 //表示和数据库的最大链接数,0表示没有限制
		IdleTimeout: 100 * time.Second, //表示最大空闲时间
		Dial: func() (redis.Conn, error) {
			//初始化链接的代码，链接到哪个ip
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
}

func main() {
	//先从pool获取一个链接
	conn := pool.Get()

	_, err := conn.Do("Set", "name", "汤姆")
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	fmt.Println("r=", r)
}
