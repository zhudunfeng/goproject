package main

import (
	"encoding/json"
	"fmt"
	_ "testing"
	"time"
)

// func TestChan(t *testing.T) {
// 	aa := make(chan interface{}, 1)
// 	aa <- "hello"
// 	close(aa)
// 	data := <-aa
// 	fmt.Println(data.(string))
// }

func Test01() {
	aa := make(chan interface{}, 1)
	aa <- "hello"
	close(aa)
	data := <-aa
	fmt.Println(data.(string))
}

func main() {
	go Test01()
	tmap := make(map[int]map[string][]int)
	tmap[0] = make(map[string][]int)
	tmap[0]["arr"] = make([]int,3)
	tmap[0]["arr"][0] = 1
	tmap[0]["arr"][1] = 2
	tmap[0]["arr"][2] = 3

	data, _ := json.Marshal(tmap)
	jsonStr := string(data)
	fmt.Println(jsonStr)
	// `{
	// 	1:{
	// 		"arr":[1,2,3]
	// 	}
	// }`

	time.Sleep(time.Second * 3)
}
