package model

import "fmt"

//定义一个结构体
type account struct {
	accountNo string
	pwd       string
	balance   float64
}

//工厂模式的函数-相当于构造函数
func NewAccount(accountNo string, pwd string, balance float64) *account {
	if len(accountNo) < 6 || len(accountNo) > 10 {
		fmt.Println("账号长度不对...")
		return nil
	}

	if len(pwd) != 6 {
		fmt.Println("密码长度不对...")
		return nil
	}

	if balance < 2 {
		fmt.Println("余额数目不对...")
		return nil
	}

	return &account{
		accountNo: accountNo,
		pwd:       pwd,
		balance:   balance,
	}
}

//getter and setter for account
func (a *account) getAccountNo() string {
	return a.accountNo
}

func (a *account) setAccountNo(accountNo string) {
	if len(accountNo) < 6 || len(accountNo) > 10 {
		fmt.Println("账号长度不对...")
		return
	}
	a.accountNo = accountNo
}

func (a *account) GetBalance() float64 {
	return a.balance
}

func (a *account) SetBalance(balance float64) {
	if balance < 2 {
		fmt.Println("余额数目不对...")
		return
	}
	a.balance = balance
}

//方法
// 1.存款
func (a *account) Deposite(money float64, pwd string) {
	//判断密码是否正确
	if pwd != a.pwd {
		fmt.Println("您输入的密码不正确！")
		return
	}

	//判断存入的金额是否合法
	if money <= 0 {
		fmt.Println("您输入的金额不合法！")
		return
	}

	a.balance += money
	fmt.Println("存款成功~~~")
}

// 2.取款
func (a *account) WithDraw(money float64, pwd string) {
	//判断密码是否正确
	if pwd != a.pwd {
		fmt.Println("您输入的密码不正确！")
		return
	}

	//判断存入的金额是否合法
	if money <= 0 {
		fmt.Println("您输入的金额不合法！")
		return
	}

	a.balance -= money
	fmt.Println("取款成功~~~")
}

// 3.查询余额
func (a *account) Query(pwd string) {
	//判断密码是否正确
	if pwd != a.pwd {
		fmt.Println("您输入的密码不正确！")
		return
	}
	fmt.Printf("您的账号为=%v 余额为=%v \n", a.accountNo, a.balance)
}
