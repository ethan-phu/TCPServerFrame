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

func (router *PingRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("call Router PreHandle")
	_, err := request.Getconnection().GetTCPConnection().Write([]byte("before ping ....\n"))
	if err != nil {
		fmt.Println("call back ping ping error", err)
	}
}
func (router *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("call Router Handle")
	_, err := request.Getconnection().GetTCPConnection().Write([]byte("now ping ....\n"))
	if err != nil {
		fmt.Println("call back ping ping error", err)
	}
}
func (router *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("call Router PostHandle")
	_, err := request.Getconnection().GetTCPConnection().Write([]byte("after ping ....\n"))
	if err != nil {
		fmt.Println("call back ping ping error", err)
	}
}
func main() {
	s := znet.NewServer("[zinx 0.1]")
	s.AddRouter(&PingRouter{})
	s.Serve()
}
