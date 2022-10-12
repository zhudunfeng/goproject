package main
 
import "fmt"

func Hello(){
	fmt.Println("hello golang")
}

//go 1.17与泛型go run -gcflags=-G=3 ./test.go
func findFunc[T comparable](a []T, v T) int {
	for i, e := range a {
		if e == v {
			return i
		}
	}
	return -1
}
func main() {
	Hello()
	fmt.Println(findFunc([]int{1, 2, 3, 4, 5, 6}, 5))
	fmt.Println(findFunc([]string{"dudu", "yiyi", "8号"}, "dudu"))
}

