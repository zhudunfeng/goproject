package main

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

//定义一个全局的pool
var pool *redis.Pool

func initPool(address string, maxIdle int, maxActive int, idleTimeout time.Duration) {
	pool = &redis.Pool{
		MaxIdle:     maxIdle,     //最大空闲数
		MaxActive:   maxActive,   //表示和数据库的最大链接数,0表示没有限制
		IdleTimeout: idleTimeout, //表示最大空闲时间
		Dial: func() (redis.Conn, error) {
			//初始化链接的代码，链接到哪个ip
			return redis.Dial("tcp", address)
		},
	}
}
