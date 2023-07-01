package process

import (
	"chat_room/client/utils"
	"chat_room/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
}

func (this *UserProcess) Register(userId int, userPwd string, userName string) (err error) {
	// fmt.Printf("usrId=%d,usrPwd=%s", usrId, usrPwd)
	// return err
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("conn err=", err)
		return
	}
	//延时关闭
	defer conn.Close()
	var mes message.Message
	mes.Type = message.RegisterType
	//创建一个LoginMes结构体
	var RegisterMes message.RegisterMes

	RegisterMes.User.UserId = userId
	RegisterMes.User.UserPwd = userPwd
	RegisterMes.User.UserName = userName

	//LoginMes结构体序列化
	data, err := json.Marshal(RegisterMes)
	if err != nil {
		fmt.Println("json err=", err)
		return
	}
	//data赋给mes
	mes.Data = string(data)
	//将mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("mes json err=", err)
		return
	}

	//建一个与服务器通信的实例
	tf := &utils.Transfer{
		Conn: conn,
	}
	//发送data给服务器端
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("发送注册信息错误=", err)
	}
	//读取服务器发送的数据
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg err=", err)
		return
	}

	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)

	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	if registerResMes.Code == 200 {
		fmt.Println("注册成功，你重新登陆一下吧")

	} else {
		fmt.Println(registerResMes.Error)

	}
	return
}
func (this *UserProcess) Login(usrId int, usrPwd string) (err error) {
	// fmt.Printf("usrId=%d,usrPwd=%s", usrId, usrPwd)
	// return err
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("conn err=", err)
		return
	}
	//延时关闭
	defer conn.Close()
	//通过conn发送数据
	var mes message.Message
	mes.Type = message.LoginMesType
	//创建一个LoginMes结构体
	var loginMes message.LoginMes

	loginMes.UserId = usrId
	loginMes.UserPwd = usrPwd
	//LoginMes结构体序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json err=", err)
		return
	}
	//data赋给mes
	mes.Data = string(data)
	//将mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("mes json err=", err)
		return
	}
	//7.data为我们要发送的数据
	//7.1 先把data的长度发送给服务器
	//先获取长度，然后将data长度转为一个切片
	var pkgLen uint32
	pkgLen = uint32(len(data))
	buffer := make([]byte, 4)

	binary.BigEndian.PutUint32(buffer, pkgLen)
	//发送长度2=
	n, err := conn.Write(buffer)
	if n != 4 || err != nil {
		fmt.Println("发送失败 err=", err)
		return
	}
	//fmt.Printf("客户端发送长度%d成功\n，内容是%s", len(data), string(data))

	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("发送失败 err=", err)
		return
	}
	//休眠20秒，
	// time.Sleep(time.Second * 5)
	// fmt.Println("休眠了20秒")
	//处理服务器端返回的数据
	tf := &utils.Transfer{
		Conn: conn,
	}
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg err=", err)
		return
	}
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)

	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	if loginResMes.Code == 200 {
		fmt.Println("登陆成功")
		//显示当前在线的用户列表
		fmt.Println("当前在线的用户列表如下")
		for _, v := range loginResMes.UserIds {
			fmt.Println("用户id:\t", v)
			user := &message.User{
				UserId:     v,
				UserStatus: message.UserOnline,
			}
			onlineUsers[v] = user
		}
		go serverProcessMes(conn)
		for {
			ShowMenu()
		}

	} else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
	}
	return
}
