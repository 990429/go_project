package process

import (
	"chat_room/common/message"
	"chat_room/server/model"
	"chat_room/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	Conn net.Conn
	//增加一个字段，表明是哪个用户的
	UserId int
}

//通知所有在线用户的方法
func (this *UserProcess) NotifyOthersOnlineUser(userId int) {
	//遍历onlineUsrs，然后一个个发送NotifyUserStatusMes
	for id, up := range userMgr.onlineUsers {
		//过滤自己
		if id == userId {
			continue
		}
		//开始通知【单独写一个方法】
		up.NotifyMeOnline(userId)
	}

}
func (this *UserProcess) NotifyMeOnline(userId int) {
	//组装我们的NotifyUserStatueMes
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType
	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOnline
	//序列化
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//序列化后信息赋值给Data
	mes.Data = string(data)
	//对mes再次序列化，准备发送
	data, err = json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//发送，创建transfer实例
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("发送数据失败")
		return
	}

}
func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {

	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json Unmarshal err=", err)
		return
	}

	var resMes message.Message //回应给客户端的数据
	resMes.Type = message.RegisterResMesType

	var registerResMes message.RegisterResMes
	//去数据库完成注册
	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = model.ERROR_USER_EXISTS.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "未知错误"
		}
	} else {
		registerResMes.Code = 200
	}
	//将loginResMes序列化
	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//4.将data赋值给resMes
	resMes.Data = string(data)

	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	//6.发送数据
	err = tf.WritePkg(data)
	return
}

//serverprocessLogin,处理登录
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	//1.先从mes中取出mes.Data，反序列化
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("1 Unmarshal err=", err)
		return
	}
	var resMes message.Message //回应给客户端的数据
	resMes.Type = message.LoginResMesType

	var loginResMes message.LoginResMes
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误"
		}

	} else {
		loginResMes.Code = 200
		//将登录成功的userId赋给this
		this.UserId = loginMes.UserId
		//用户登录成功，将其放入userMgr中
		userMgr.AddOnlineUser(this)
		//通知其他用户该用户上线
		this.NotifyOthersOnlineUser(loginMes.UserId)
		//将当前用户id放入loginresmes.id
		//遍历userMgr.onlineUsers

		for id := range userMgr.onlineUsers {
			loginResMes.UserIds = append(loginResMes.UserIds, id)
		}
		fmt.Println(user, "登录成功！")

	}

	fmt.Println("用户发送信息为：", loginMes)
	//将loginResMes序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//4.将data赋值给resMes
	resMes.Data = string(data)

	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	//6.发送数据
	err = tf.WritePkg(data)
	return
}
