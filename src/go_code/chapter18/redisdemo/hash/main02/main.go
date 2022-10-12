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
		fmt.Println("redis.Dial() error: ", err)
		return
	}

	//延迟关闭，回收资源
	defer conn.Close()

	//2.通过go 向redis写入数据 hash
	_, err = conn.Do("Hmset", "user02", "name", "Tom", "age", 18)
	if err != nil {
		fmt.Println("Hmset error: ", err)
		return
	}

	//3.通过go向redis读取数据 hash
	r, err := redis.Strings(conn.Do("Hmget", "user02", "name", "age"))
	if err != nil {
		fmt.Println("Hmget error: ", err)
		return
	}

	//因为返回r是interface{}
	//因为name 对应的值是string，因此我们需要转换

	for i, v := range r {
		fmt.Printf("r[%d]=%v\n", i, v)
	}

}
