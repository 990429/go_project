package main

import (
	"chat_room/client/process"
	"fmt"
)

var userId int
var userPwd string
var userName string

func main() {
	var key int
	var loop = true
	for loop {
		fmt.Println("------------欢迎登陆多人聊天系统-------------")
		fmt.Println("\t\t\t 1 登陆聊天系统")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 请选择（1-3）")
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登陆聊天室")
			fmt.Println("请输入你的id号")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("输入你的密码")
			fmt.Scanf("%s\n", &userPwd)
			up := &process.UserProcess{}
			up.Login(userId, userPwd)

		case 2:
			fmt.Println("注册用户")
			fmt.Println("请输入注册的id")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户的密码")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("请输入用户的名字（nickname）")
			fmt.Scanf("%s\n", &userName)
			//创建一个user实例
			up := &process.UserProcess{}
			up.Register(userId, userPwd, userName)
		case 3:
			fmt.Println("退出系统")
			loop = false
		default:
			fmt.Println("输入有误，重新输入")
		}
	}

}
