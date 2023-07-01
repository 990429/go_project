package utils

import (
	"chat_room/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

//将这些方法关联到结构体中
type Transfer struct {
	Conn   net.Conn
	Buffer [8096]byte
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	//buffer := make([]byte, 8096)
	fmt.Println("读取服务器发送数据")
	_, err = this.Conn.Read(this.Buffer[:4])
	if err != nil {
		fmt.Println("conn.Read err=", err)
		//err = errors.New("read pkg head err")
		return
	}
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buffer[:4])
	//根据pkgLen读取消息内容
	n, err := this.Conn.Read(this.Buffer[:pkgLen])
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn read err=", err)
		return
	}

	//把pkglen反序列化为messag.Message
	err = json.Unmarshal(this.Buffer[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
	}
	return
}

//编写函数发送数据writePkg

func (this *Transfer) WritePkg(data []byte) (err error) {
	//先发送一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	//buffer := make([]byte, 4)

	binary.BigEndian.PutUint32(this.Buffer[:4], pkgLen)
	//发送长度2=
	n, err := this.Conn.Write(this.Buffer[:4])
	if n != 4 || err != nil {
		fmt.Println("发送失败 err=", err)
		return
	}
	_, err = this.Conn.Write(data)
	if err != nil {
		fmt.Println("发送失败 err=", err)
		return
	}
	fmt.Printf("客户端端发送长度%d成功\n，内容是%s\n", len(data), string(data))
	return
}
