package main

import "fmt"

//定义一个结构体
type Account struct {
	AccountNo string
	Pwd       string
	Balance   float64
}

//方法
// 1.存款
func (a *Account) Deposite(money float64, pwd string) {
	//判断密码是否正确
	if pwd != a.Pwd {
		fmt.Println("您输入的密码不正确！")
		return
	}

	//判断存入的金额是否合法
	if money <= 0 {
		fmt.Println("您输入的金额不合法！")
		return
	}

	a.Balance += money
	fmt.Println("存款成功~~~")
}

// 2.取款
func (a *Account) WithDraw(money float64, pwd string) {
	//判断密码是否正确
	if pwd != a.Pwd {
		fmt.Println("您输入的密码不正确！")
		return
	}

	//判断存入的金额是否合法
	if money <= 0 {
		fmt.Println("您输入的金额不合法！")
		return
	}

	a.Balance -= money
	fmt.Println("取款成功~~~")
}

// 3.查询余额
func (a *Account) Query(pwd string) {
	//判断密码是否正确
	if pwd != a.Pwd {
		fmt.Println("您输入的密码不正确！")
		return
	}
	fmt.Printf("您的账号为=%v 余额为=%v \n", a.AccountNo, a.Balance)
}

func main() {
	account := Account{
		AccountNo: "ADUN521",
		Pwd:       "666666",
		Balance:   100,
	}

	menuCode := 0

	for {
		fmt.Printf("\n\n====================\n")
		fmt.Println("请选择菜单操作项：")
		fmt.Println("1.存款")
		fmt.Println("2.取款")
		fmt.Println("3.查询余额")
		fmt.Println("4.退出")
		fmt.Printf("====================\n\n")

		fmt.Scanln(&menuCode)

		switch menuCode {
		case 1:
			fmt.Println("请输入密码：")
			pwd := ""
			fmt.Scanln(&pwd)
			fmt.Println("请输入存款金额：")
			money := 0.0
			fmt.Scanln(&money)
			account.Deposite(money, pwd)
		case 2:
			fmt.Println("请输入密码：")
			pwd := ""
			fmt.Scanln(&pwd)
			fmt.Println("请输入取款金额：")
			money := 0.0
			fmt.Scanln(&money)
			account.WithDraw(money, pwd)
		case 3:
			fmt.Println("请输入密码：")
			pwd := ""
			fmt.Scanln(&pwd)
			account.Query(pwd)
		case 4:
			return
		default:
			fmt.Println("菜单项输入有误，请重新输入")
		}
	}

}
