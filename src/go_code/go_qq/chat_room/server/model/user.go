package model

//定义一个结构体
type User struct {
	UserId   int    `json:"usrId"`
	UserPwd  string `json:"usrPwd"`
	UserName string `json:"usrName"`
}
