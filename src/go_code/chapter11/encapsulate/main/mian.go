package main

import (
	"fmt"
	"go_code/chapter11/encapsulate/model"
)

func main() {
	var person = model.NewPerson("Tom")
	person.SetAge(18)
	person.SetSal(25000)
	fmt.Printf("name = %v age = %v sal=%v\n",
		person.Name, person.GetAge(), person.GetSal())
}
