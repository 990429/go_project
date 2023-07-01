package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func main() {

	//fmt.Printf("客户端总共发送字节数据\n")
	cn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	defer cn.Close()
	_, err = cn.Do("SET", "name", "flb")
	r, err := redis.String(cn.Do("Get", "name"))

	fmt.Println(r)
	//fmt.Printf("sucess %v", cn)
}
