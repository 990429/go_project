package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func receive(conn net.Conn) {
	fmt.Println("receive")
	buf := make([]byte, 1024)
	for {

		//创建一个新的切片

		//等待，直到客户端通过conn发送信息，
		//如果客户端没有write发送，阻塞
		fmt.Println("conn.Read(buf)")
		n, err := conn.Read(buf)

		if err != nil {
			fmt.Println("客户端read err=", err)
			return
		}
		data := string(buf[:n])
		fmt.Printf(data)

	}
}
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("client dial err=", err)
		return
	}
	fmt.Println("conn suc,", conn)
	reader := bufio.NewReader(os.Stdin)
	count := 0
	go receive(conn)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("readerstring err=", err)
		}
		line_ := strings.Trim(line, " \r\n")
		if line_ == "exit" {
			fmt.Println("客户端退出")
			break
		}
		n, err := conn.Write([]byte(line))
		if err != nil {
			fmt.Println("conn.write err=", err)
		}
		count += n
	}
	fmt.Printf("客户端总共发送%d字节数据\n", count)
}
