package main

import (
	"fmt"
	"go_code/chapter13/customerManager/model"
	"go_code/chapter13/customerManager/service"
)

type CustomerView struct {
	//定义必要的字段
	key  string //接收用户输入...
	loop bool   //表示是否循环的显示主菜单
	//增加一个字段customerService
	customerService *service.CustomerService
}

//显示所有的客户信息
func (this *CustomerView) list() {
	//首先，获取到当前所有的客户信息(在切片中)
	customers := this.customerService.List()
	//显示
	fmt.Println("---------------------------客户列表---------------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t\t邮箱")
	for _, customer := range customers {
		fmt.Println(customer.GetInfo())
	}
	fmt.Printf("\n-------------------------客户列表完成-------------------------\n\n")
}

//得到用户的输入，信息构建新的客户，并完成添加
func (this *CustomerView) add() {
	fmt.Println("---------------------添加客户---------------------")
	fmt.Println("姓名:")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别:")
	sex := ""
	fmt.Scanln(&sex)
	fmt.Println("年龄:")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话:")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱:")
	email := ""
	fmt.Scanln(&email)
	//构建一个新的Customer实例
	//注意: id号，没有让用户输入，id是唯一的，需要系统分配
	customer := model.NewCustomer2(name, sex, age, phone, email)
	//调用
	if this.customerService.Add(customer) {
		fmt.Println("---------------------添加完成---------------------")
	} else {
		fmt.Println("---------------------添加失败---------------------")
	}
}

//得到用户的输入id，删除该id对应的客户
func (this *CustomerView) delete() {
	fmt.Println("---------------------删除客户---------------------")
	fmt.Println("请选择待删除客户编号(-1退出)：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return //放弃删除操作
	}
	choice := ""
	for {
		fmt.Println("确认是否删除(Y/N)：")
		//这里同学们可以加入一个循环判断，直到用户输入 y 或者 n,才退出..
		fmt.Scanln(&choice)
		if choice == "y" || choice == "Y" || choice == "n" || choice == "N" {
			break
		}
	}
	if choice == "y" || choice == "Y" {
		//调用customerService 的 Delete方法
		if this.customerService.Delete(id) {
			fmt.Println("---------------------删除完成---------------------")
		} else {
			fmt.Println("--------------删除失败，输入的id号不存在------------")
		}
	}

}

//得到用户的输入，信息构建新的客户，并完成添加
func (this *CustomerView) update() {
	fmt.Println("---------------------修改客户---------------------")
	fmt.Println("请选择待修改客户编号(-1退出)：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return //放弃删除操作
	}
	oriCustomer, _ := this.customerService.FindById(id)
	fmt.Printf("姓名(%v):", oriCustomer.Name)
	name := ""
	fmt.Scanln(&name)
	fmt.Printf("性别(%v):", oriCustomer.Gender)
	sex := ""
	fmt.Scanln(&sex)
	fmt.Printf("年龄(%v):", oriCustomer.Age)
	age := 0
	fmt.Scanln(&age)
	fmt.Printf("电话(%v):", oriCustomer.Phone)
	phone := ""
	fmt.Scanln(&phone)
	fmt.Printf("邮箱(%v):", oriCustomer.Email)
	email := ""
	fmt.Scanln(&email)
	//构建一个新的Customer实例
	//注意: id号，没有让用户输入，id是唯一的，需要系统分配
	customer := model.NewCustomer2(name, sex, age, phone, email)
	customer.Id = id
	//调用
	if this.customerService.Update(customer) {
		fmt.Println("---------------------修改完成---------------------")
	} else {
		fmt.Println("---------------------修改失败---------------------")
	}
}

//退出软件
func (this *CustomerView) exit() {
	fmt.Println("确认是否退出(Y/N)：")
	for {
		fmt.Scanln(&this.key)
		if this.key == "y" || this.key == "Y" || this.key == "n" || this.key == "N" {
			break
		}
		fmt.Println("你的输入有误，确认是否退出(Y/N)：")
	}
	if this.key == "y" || this.key == "Y" {
		this.loop = false
	}
}

//显示主菜单
func (this *CustomerView) mainMenu() {
	for {
		fmt.Println("-----------------客户信息管理软件-----------------")
		fmt.Println("                 1 添 加 客 户")
		fmt.Println("                 2 修 改 客 户")
		fmt.Println("                 3 删 除 客 户")
		fmt.Println("                 4 客 户 列 表")
		fmt.Println("                 5 退       出")
		fmt.Println("-------------------------------------------------")
		fmt.Print("请选择(1-5)：")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.add()
		case "2":
			this.update()
		case "3":
			this.delete()
		case "4":
			this.list()
		case "5":
			this.exit()
		default:
			fmt.Println("你的输入有误，请重新输入...")
		}

		if !this.loop {
			break
		}
	}
	fmt.Println("你退出了客户关系管理系统...")
}

func main() {
	customerView := CustomerView{
		key:  "",
		loop: true,
	}
	//这里完成对customerView结构体的customerService字段的初始化
	customerView.customerService = service.NewCustomerService()

	customerView.mainMenu()
}
