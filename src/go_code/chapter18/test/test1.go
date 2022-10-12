//go 1.17与泛型go run -gcflags=-G=3 ./test.go

package main

import "fmt"

// Index returns the index of x in s, or -1 if not found.
func haha[T comparable](s []T, x T) int {
 for i, v := range s {
  // v and x are type T, which has the comparable
  // constraint, so we can use == here.\
	if v==x{
		fmt.Println(v)
		return i
	}
 }
 return -1
}

func main() {
 // Index works on a slice of ints

  // Index also works on a slice of strings
 ss := []string{"foo", "bar", "baz"}
 fmt.Println(haha(ss, "hello"))
}