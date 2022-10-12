package apply05

import (
	"reflect"
	"testing"
)

type user struct {
	UserId string
	Name   string
}

// 使用反射创建并操作结构体
func TestReflectStructPtr(t *testing.T) {
	var (
		model *user
		st    reflect.Type
		elem  reflect.Value
	)

	st = reflect.TypeOf(model)                                 //获取类型*user
	t.Log("reflect.TypeOf(model):", st.Kind().String())        //ptr
	st = st.Elem()                                             //st指向的类型
	t.Log("reflect.TypeOf(model).Elem():", st.Kind().String()) //struct
	elem = reflect.New(st)                                     //New返回一个Value类型值，该值持有一个指向类型为type的新申请的零值的指针
	t.Log("reflect.New:", elem.Kind().String())                //ptr
	t.Log("reflect.New.Elem", elem.Elem().Kind().String())     //struct
	//model就是创建的user结构体变量（实例）
	model = elem.Interface().(*user) //model是*user它的指向和elem是一样的

	elem = elem.Elem()                           //取得elem指向的值【指针指向的具体】
	elem.FieldByName("UserId").SetString("1111") //赋值
	elem.FieldByName("Name").SetString("nickname")

	t.Log("model.Name", model.Name)

}
