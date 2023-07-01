package main

import (
	"chat_room/server/model"
	"fmt"
	"net"
	"time"
)

func process_(conn net.Conn) {
	//读取客户端发送的信息
	//需要延时关闭
	//将读取数据包封装成一个函数
	defer conn.Close()
	processor := &Processor{
		Conn: conn,
	}
	err := processor.main_processor()
	if err != nil {
		fmt.Println("通讯间协程出问题")
	}
}

func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool)
}
func init() {
	//初始化连接池
	initPool("localhost:6379", 16, 0, 300*time.Second)
	initUserDao()
}
func main() {

	//提示信息
	fmt.Println("服务器【新的结构】在8889端口监听")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("net listen err=", err)
		return
	}
	for {

		fmt.Println("等待客户端来连接服务器")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen acc err=", err)
			continue //如果连接失败则返回循环头
		}
		//连接成功后，启动协程保持与客户端的通讯
		go process_(conn)
	}
}
