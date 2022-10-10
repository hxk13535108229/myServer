/*
 * @Date: 2022-10-10 18:24:54
 * @LastEditors: hxk
 * @LastEditTime: 2022-10-10 18:46:00
 * @FilePath: \myServer\testDemo\xuebao_server\server.go
 */
package main

import (
	"fmt"
	"xuebao/kiface"
	"xuebao/knet"
)

//ping test 自定义路由
type PingRouter struct {
	knet.BaseRouter //一定要先基础BaseRouter
}

//Test PreHandle
func (this *PingRouter) PreHandle(request kiface.IRequest) {
	fmt.Println("Call Router PreHandle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping ....\n"))
	if err != nil {
		fmt.Println("call back ping ping ping error")
	}
}

//Test Handle
func (this *PingRouter) Handle(request kiface.IRequest) {
	fmt.Println("Call PingRouter Handle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping...ping...ping\n"))
	if err != nil {
		fmt.Println("call back ping ping ping error")
	}
}

//Test PostHandle
func (this *PingRouter) PostHandle(request kiface.IRequest) {
	fmt.Println("Call Router PostHandle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("After ping .....\n"))
	if err != nil {
		fmt.Println("call back ping ping ping error")
	}
}

func main() {
	//创建一个server句柄
	s := knet.NewServer("xuebao server")

	s.AddRouter(&PingRouter{})

	//2 开启服务
	s.Serve()
}
