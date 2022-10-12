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
	_, err = conn.Do("Hset", "user01", "name", "Tom")
	if err != nil {
		fmt.Println("Hset error: ", err)
		return
	}

	_, err = conn.Do("Hset", "user01", "age", 18)
	if err != nil {
		fmt.Println("Hset error: ", err)
		return
	}

	//3.通过go向redis读取数据 hash
	r1, err := redis.String(conn.Do("Hget", "user01", "name"))
	if err != nil {
		fmt.Println("Hget error: ", err)
		return
	}

	r2, err := redis.Int(conn.Do("Hget", "user01", "age"))
	if err != nil {
		fmt.Println("Hget error: ", err)
		return
	}

	//因为返回r是interface{}
	//因为name 对应的值是string，因此我们需要转换
	fmt.Printf("result r1: %v \t r2: %v \n", r1, r2)

}
