package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func main() {
	//通过go向redis写入数据和读取数据
	//1.连接到redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Printf("redis.Dial err:%v\n", err)
		return
	}

	defer conn.Close() //关闭，回收资源

	//2.通过go 向redis写入数据 string [key-value]
	_, err = conn.Do("Set", "name", "Tom")
	if err != nil {
		fmt.Println("set err:", err)
		return
	}

	//3.通过go向redis读取数据 string [key-value]
	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("get err:", err)
		return
	}
	//因为返回r是interface{}
	//因为name 对应的值是string，因此我们需要转换
	fmt.Println("get ok:", r)
}
