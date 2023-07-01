package model

import (
	"chat_room/common/message"
	"encoding/json"
	"fmt"

	"github.com/gomodule/redigo/redis"
)

//定义一个userDao结构体
//完成对user的各种处理

type UserDao struct {
	pool *redis.Pool
}

var (
	MyUserDao *UserDao
)

//使用工厂模式，创建一个UserDao实例
func NewUserDao(pool *redis.Pool) (userdao *UserDao) {
	userdao = &UserDao{
		pool: pool,
	}
	return
}

//1.根据用户id返回一个user实例+err
func (this *UserDao) GetUserById(conn redis.Conn, id int) (user User, err error) {
	//通过给定id去redis里查询
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		if err == redis.ErrNil {
			err = ERROR_USER_NOTEXISTS
		}
	}
	//user = &User{}
	//把res进行反序列化·1
	fmt.Println("res=", len(res))
	//fmt.Println("res_byte=", []byte(res))
	res_byte := []byte(res)
	err = json.Unmarshal(res_byte, &user)
	if err != nil {
		fmt.Println("userDao json.Unmarshal err=", err)
		return
	}
	return
}

//完成对用户的校验 Login
//1.login完成对用户的验证
//2.如果id和pwd均正确，返回一个user实例
//3.如果有误，返回err
func (this *UserDao) Login(UserId int, UserPwd string) (user User, err error) {
	//先从连接池中取出一个连接
	conn := this.pool.Get()
	defer conn.Close()
	user, err = this.GetUserById(conn, UserId)
	if err != nil {
		return
	}
	//这时证明这个用户获取到
	if user.UserPwd != UserPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}

func (this *UserDao) Register(user *message.User) (err error) {
	//先从连接池中取出一个连接
	conn := this.pool.Get()
	defer conn.Close()
	_, err = this.GetUserById(conn, user.UserId)
	if err == nil {
		err = ERROR_USER_EXISTS
		return
	}
	//这时证明这个用户还没有注册过
	data, err := json.Marshal(user)
	if err != nil {
		return
	}
	//入库
	_, err = conn.Do("HSet", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("用户注册失败，err=", err)
		return
	}
	return
}
