package process

import (
	"chat_room/common/message"
	"chat_room/server/utils"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

//1.显示登陆成功的界面
func ShowMenu() {
	fmt.Println("------恭喜xxx登陆成功--------")
	fmt.Println("------1.显示在线用户列表-----")
	fmt.Println("------2.发送消息-------------")
	fmt.Println("------3.信息列表------------")
	fmt.Println("------4.退出系统------------")
	fmt.Println("请选择1-4：")
	var key int
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		fmt.Println("显示在线用户列表")
		outputOnlineUser()
	case 2:
		fmt.Println("发送消息")
	case 3:
		fmt.Println("消息列表")
	case 4:
		fmt.Println("你选择了退出系统")
		os.Exit(0)
	default:
		fmt.Println("你输入的不对")
	}
}

///和服务器保持通讯
func serverProcessMes(conn net.Conn) {
	//1.创建一个transfer实例，不停读取服务器发送的数据
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		fmt.Println("客户端%s正在等待读取服务器发送的数据")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg err=", err)
			return
		}
		//fmt.Println("mes=%v", mes)
		switch mes.Type {
		case message.NotifyUserStatusMesType: //有人上线了
			var notifyUserStatusMes message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			updateUserStatus(&notifyUserStatusMes)
		//处理
		//1.取出状态
		//2.把这个用户的信息，状态保存到客户map中

		default:
			fmt.Println("服务器端返回了未知的消息类型")
		}
	}
}
