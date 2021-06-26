package znet

import (
	"errors"
	"fmt"
	"net"
	"zinx/ziface"
)

// Iserver的接口实现，定义一个Server的服务模块
type Server struct {
	// 服务器名称
	Name string
	// 服务器编订的Ip版本
	IPVersion string
	// 服务器监听的Ip
	IP string
	// 服务器监听的端口
	Port int
}

/*
创建一个服务器句柄
*/
func NewServer(name string) ziface.Iserver {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}
	return s
}

// 定义当前客户端链接的handle api
func CallBackToClient(conn *net.TCPConn, data []byte, cnt int) error {
	//回显业务
	fmt.Println("[Conn Handle] CallBackToClient...")
	_, err := conn.Write(data[:cnt])
	if err != nil {
		fmt.Println("write back buf err", err)
		return errors.New("CallBackToClient error")
	}
	return nil
}

// 开启网络服务
func (s *Server) Start() {
	fmt.Printf("[Start] Server Listenner at Ip:%s ,Port :%d is starting\n", s.IP, s.Port)
	go func() {
		// 1 获取一个TCP的Addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr error:", err)
			return
		}
		// 2 监听服务器的地址
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen", s.IPVersion, "err", err)
		}
		fmt.Println("start Zinx server succ,", s.Name, "succ,Listenning....")
		// 3 阻塞的等待客户端连接，处理客户端连接业务
		// TODO server.go 应该应该有一个自动生成ID的方法
		var cid uint32
		cid = 0
		for {
			// 如果有客户端链接过来，阻塞会返回
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err", err)
				continue
			}
			//已经与客户端建立连接，做一些业务，做一个最基本的512字节长度的回写业务
			clientConn := NewConnection(conn, cid, CallBackToClient)
			cid++
			// 3.4 启动当前链接处理的业务
			go clientConn.Start()
		}
	}()
}

func (s *Server) Stop() {
	// TODO 将一些服务器的资源，状态，或者一些已经开辟的连接信息 进行停止或者回收
	fmt.Println("[STOP] zinx server,name", s.Name)
}
func (s *Server) Serve() {
	// 启动server的服务功能
	s.Start()
	// TODO 做一些启动服务器之后的额外业务
	// 阻塞状态
	select {}
}
