package apploy04

import (
	"reflect"
	"testing"
)

type User struct {
	UserId string
	Name   string
}

// 4)使用反射操作任意结构体类型:【了解】
func TestReflectStruct(t *testing.T) {
	var (
		model *User
		sv    reflect.Value
	)

	model = &User{}
	sv = reflect.ValueOf(model)
	t.Log("reflect.Value", sv.Kind().String())        //ptr
	sv = sv.Elem()                                    //获取指针指向的值
	t.Log("reflect.ValueOf.Elem", sv.Kind().String()) //struct
	sv.FieldByName("UserId").SetString("12345")
	sv.FieldByName("Name").SetString("Jack")
	t.Log("model:", *model)
}
