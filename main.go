package main

import (
	"fmt"
	"zinx/ziface"
	"zinx/znet"
)

// ping test 自定义路由
type PingRouter struct {
	znet.BaseRouter // 一定要先基础BaseRouter
}

func (router *PingRouter) Handle(request ziface.IRequest) {
	// 先读取客户端发送过来的数据
	fmt.Println("Call PingRouter Handle")
	//先读取客户端的数据，再回写ping...ping...ping
	fmt.Println("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))
	fmt.Println("call Router PreHandle")
	err := request.Getconnection().SendMsg(0, []byte("before ping ping ping"))
	if err != nil {
		fmt.Println("call back ping ping error", err)
	}
}

// 构建 pong router 自定义路由
type PongRouter struct {
	znet.BaseRouter // 一定要先基础路由BaseRouter
}

func (router *PongRouter) Handle(request ziface.IRequest) {
	// 先读取客户端发送过来的数据
	fmt.Println("Call PingRouter Handle")
	//先读取客户端的数据，再回写ping...ping...ping
	fmt.Println("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))
	fmt.Println("call Router PreHandle")
	err := request.Getconnection().SendMsg(1, []byte("before pong pong pong"))
	if err != nil {
		fmt.Println("call back ping ping error", err)
	}
}
func main() {
	s := znet.NewServer()
	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &PongRouter{})
	s.Serve()
}
