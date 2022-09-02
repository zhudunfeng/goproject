package testdemo1

import (
	_ "fmt"
	"testing"
)

//编写要给测试用例,去测试addUpper是否正确
func TestAddUpper(t *testing.T) {
	//调用
	res := addUpper(10)
	if res != 55 {
		// fmt.Printf("AddUpper(10)执行错误，期望值=%v 实际值=%v\n", 55, res)
		t.Fatalf("AddUpper(10)执行错误，期望值=%v 实际值=%v\n", 55, res)
	}
	t.Logf("AddUpper(10)执行正确...")
}

func TestHello(t *testing.T) {
	//调用
	Hello()
}
