package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	//循环接收客户端发送数据
	defer conn.Close()
	fmt.Printf("等待客户端%s 发送数据\n", conn.RemoteAddr().String())
	buf := make([]byte, 1024)
	for {

		//创建一个新的切片

		//等待，直到客户端通过conn发送信息，
		//如果客户端没有write发送，阻塞
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("服务器端read err=", err)
			return
		}
		data := string(buf[:n])
		if data == "exit\n" {
			fmt.Println("客户端自动退出")
			break
		}
		fmt.Printf(data)
		conn.Write([]byte("服务器接收成功！"))
	}
}
func main() {
	fmt.Println("服务器开始监听")
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("listen err", err)
		return
	}
	defer listen.Close()

	for {
		fmt.Println("等待客户端进行连接。。。")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept() err=", err)
		} else {
			fmt.Printf("accept() suc con=%v, 客户端ip=%v\n", conn, conn.RemoteAddr().String())
		}
		go process(conn)
	}
}
