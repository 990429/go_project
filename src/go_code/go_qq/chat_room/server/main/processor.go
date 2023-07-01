package main

import (
	"chat_room/common/message"
	"chat_room/server/process"
	"chat_room/server/utils"
	"fmt"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

///serverprocessMes函数
//根据客户端发送的消息类型的不同，调用不用函数来处理
func (this *Processor) serverprocessMes(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		up := &process.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterType:
		up := &process.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessRegister(mes)
	default:
		fmt.Println("消息类型不存在，无法处理")
	}
	return
}
func (this *Processor) main_processor() (err error) {
	for {
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("对方关闭了连接")
			} else {
				fmt.Println("readPkg err=", err)
			}
			return err
		}
		fmt.Printf("用户发送数据为%v\n", mes.Data)
		err = this.serverprocessMes(&mes)
		if err != nil {
			fmt.Println("serverProcess err=", err)
			return err
		}
		//fmt.Println("mes=", mes)
	}
}
