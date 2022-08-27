package main

import "fmt"

/*
1)一个景区根据游人的年龄收取不同价格的门票，比如年龄大于18，收费20元，其它情况门票免费.
2)请编写Visitor结构体，根据年龄段决定能够购买的门票价格并输出t
*/

type Visitor struct {
	Name string
	Age  int
}

func (visitor *Visitor) showPrice() {
	if visitor.Age < 8 || visitor.Age > 90 {
		fmt.Println("考虑到安全，请不要游玩")
		return
	}

	if visitor.Age > 18 {
		fmt.Printf("游客的名字为name=[%v] age=[%v] 收费为20\n", visitor.Name, visitor.Age)
	} else {
		fmt.Printf("游客的名字为name=[%v] age=[%v] 免费\n", visitor.Name, visitor.Age)
	}
}
