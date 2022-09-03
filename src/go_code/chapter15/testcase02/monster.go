package monster

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

//给Monster绑定方法store，可以将一个Monster变量(对象),序列化后保存到文件中
func (m *Monster) Store() bool {
	//先序列化
	data, err := json.Marshal(m)
	if err != nil {
		fmt.Println("Marshal error:", err)
		return false
	}
	//保存到文件
	filepath := "./monster.json"
	err = ioutil.WriteFile(filepath, data, 777)
	if err != nil {
		fmt.Println("WriteFile error:", err)
		return false
	}
	return true
}

//给Monster绑定方法Restore，可以将一个序列化的Monster,从文件中读取,
//并反序列化为Monster对象,检查反序列化,名字正确
func (m *Monster) Restore() bool {
	//1．先从文件中，读取序列化的字符串
	filepath := "./monster.json"
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("ReadFile error:", err)
		return false
	}
	//2.使用读取到data []byte ,对反序列化
	err = json.Unmarshal(data, m)
	if err != nil {
		fmt.Println("Unmarshal error:", err)
		return false
	}
	return true
}
