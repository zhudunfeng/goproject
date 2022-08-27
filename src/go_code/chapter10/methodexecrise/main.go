package main

import "fmt"

type MethondUtils struct {
	//字段
}

//给MethondUtils编写方法
func (m *MethondUtils) Print() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

}

//2) 编写一个方法，提供m和n两个参数，方法中打印一个m*n的矩形
func (mu *MethondUtils) Print2(m int, n int) {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

/*
	编写一个方法算该矩形的面积(可以接收长len，和宽width),
	将其作为方法返回值。在main方法中调用该方法,接收返回的面积值并打印
*/
func (mu *MethondUtils) area(len float64, width float64) float64 {
	return len * width
}

//4)编写方法:判断一个数是奇数还是偶数
func (m *MethondUtils) JudgeNum(num int) {
	if num%2 == 0 {
		fmt.Printf("%v是偶数\n", num)
	} else {
		fmt.Printf("%v是奇数\n", num)
	}
}

//5)根据行、列、字符打印对应行数和列数的字符，比如:行:3，列:2，字符*,则打印相应的效果
func (mu *MethondUtils) Print3(m int, n int, key string) {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(key)
		}
		fmt.Println()
	}
}

// 在MethodUtils结构体编个方法,从键盘接收整数(1-9),打印对应乘法表:
func (mu *MethondUtils) JiuJiu() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v*%v=%v\t", j, i, i*j)
		}
		fmt.Println()
	}
}

/*
	定义小小计算器结构体(Calcuator) ,
	实现加减乘除四个功能
	实现形式1:分四个方法完成:,分别计算+-
	实现形式2:用一个方法搞定,需要接收两个数,还有一个运算符
*/
type Calcuator struct {
	Num1 float64
	Num2 float64
}

//实现形式1
func (c *Calcuator) add() float64 {
	return c.Num1 + c.Num2
}

func (c *Calcuator) sub() float64 {
	return c.Num1 - c.Num2
}

//实现形式2
func (c *Calcuator) getRes(operator byte) float64 {
	res := 0.0
	switch operator {
	case '+':
		res = c.Num1 + c.Num2
	case '-':
		res = c.Num1 - c.Num2
	case '*':
		res = c.Num1 * c.Num2
	case '/':
		res = c.Num1 / c.Num2
	default:
		fmt.Println("操作符输入有误！")
	}
	return res
}

func (c *Calcuator) swap(arr *[3][3]int) {
	for i := 0; i < len(*arr); i++ {
		for j := 0; j < i; j++ {
			//交换
			temp := (*arr)[i][j]
			(*arr)[i][j] = (*arr)[j][i]
			(*arr)[j][i] = temp
		}
	}
}

func (c *Calcuator) swap1(arr [][]int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			//交换
			temp := arr[i][j]
			arr[i][j] = arr[j][i]
			arr[j][i] = temp
		}
	}
}

func main() {
	/*
		1)编写结构体(Methodutils)，编程一个方法,方法不需要参数,
		在方法中打印一个10*8的矩形,在main方法中调用该方法。
	*/
	var m MethondUtils
	m.Print()
	fmt.Println()
	m.Print2(2, 6)
	fmt.Println()
	areaRes := m.area(3.0, 4.1)
	fmt.Println("areaRes=", areaRes)

	m.JudgeNum(10)

	m.Print3(5, 10, "@")

	m.JiuJiu()

	c := Calcuator{}
	c.Num1 = 1.2
	c.Num2 = 2.2
	//方式1
	add := c.add()
	sub := c.sub()
	fmt.Printf("c.add(): %.2f\n", add)
	fmt.Printf("c.sub(): %.2f\n", sub)
	fmt.Printf("add=%v\n", fmt.Sprintf("%.2f", add))
	fmt.Printf("sub=%v\n", fmt.Sprintf("%.2f", sub))

	//方式2
	fmt.Printf("c.getRes('+'): %.2f\n", c.getRes('+'))
	fmt.Printf("c.getRes('-'): %.2f\n", c.getRes('-'))
	fmt.Printf("c.getRes('*'): %.2f\n", c.getRes('*'))
	fmt.Printf("c.getRes('/'): %.2f\n", c.getRes('/'))

	//交换数组
	arr := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	// c.swap(&arr)
	c.swap1(arr)
	for i, v1 := range arr {
		for j, v2 := range v1 {
			fmt.Printf("arr[%v][%v]=%v ", i, j, v2)
		}
		fmt.Println()
	}

}
