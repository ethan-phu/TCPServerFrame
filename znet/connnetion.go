package znet

import (
	"errors"
	"fmt"
	"io"
	"net"
	"zinx/ziface"
)

/*
链接模块
*/
type Connection struct {
	// 当前链接的socket TCP套接字
	Conn *net.TCPConn
	// 链接ID
	ConnID uint32
	// 当前的链接状态
	isClosed bool
	// 该链接的处理方法router
	Router ziface.IRouter
	// 告知当前链接已经退出或停止
	ExitBuffChan chan bool
}

// 初始化链接模块方法
func NewConnection(conn *net.TCPConn, connID uint32, router ziface.IRouter) ziface.Iconnetction {
	c := &Connection{
		Conn:         conn,
		ConnID:       connID,
		isClosed:     false,
		Router:       router,
		ExitBuffChan: make(chan bool, 1),
	}
	return c
}

// 处理conn读取数据的Goroutine
func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is runing")
	defer fmt.Println(c.RemoteAddr().String(), "conn reader exit")
	defer c.Stop()
	for {
		// 创建封包拆包的对象
		dp := NewDataPack()
		// 读取我们最大的数据到buff中
		headData := make([]byte, dp.GetHeadLen())
		_, err := io.ReadFull(c.GetTCPConnection(), headData)
		if err != nil {
			fmt.Println("read msg head error:", err)
			c.ExitBuffChan <- true
			continue
		}
		// 拆包，得到msgid和datalen放在msg中
		msg, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("unpack error:", err)
			c.ExitBuffChan <- true
			continue
		}
		// 根据dataLen读取data,凡在msg.Data中
		var data []byte
		if msg.GetDataLen() > 0 {
			data = make([]byte, msg.GetDataLen())
			_, err := io.ReadFull(c.GetTCPConnection(), data)
			if err != nil {
				fmt.Println("unpack error:", err)
				c.ExitBuffChan <- true
				continue
			}
		}
		msg.SetData(data)
		// 得到当前客户端请求的Request数据
		req := Request{
			conn: c,
			msg:  msg,
		}
		// 从路由Routers中找到注册绑定Conn的对应Handle
		go func(request ziface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)
	}
}

// 启动连接，让当前的链接准备开始工作
func (conn *Connection) Start() {
	// 开启处理该链接读取到客户端数据之后的请求业务
	go conn.StartReader()
	for { //这里需要再进行优化
		select {
		case <-conn.ExitBuffChan:
			//得到退出消息，不在阻塞
			return
		}
	}
}

// 停止链接 结束当前链接的工作
func (conn *Connection) Stop() {
	if conn.isClosed {
		return
	}
	conn.isClosed = true
	// 关闭socket链接
	conn.Conn.Close()
	// 通知从缓冲队列读取数据的业务，该链接已经关闭
	conn.ExitBuffChan <- true
	// 关闭该链接的全部管道
	close(conn.ExitBuffChan)
}

// 获取当前链接的绑定socket conn
func (conn *Connection) GetTCPConnection() *net.TCPConn {
	return conn.Conn
}

// 获取当前链接模块的链接ID
func (conn *Connection) GetConnID() uint32 {
	return conn.ConnID
}

// 获取远程客户端的TCP状态IP port
func (conn *Connection) RemoteAddr() net.Addr {
	return conn.Conn.RemoteAddr()
}

// 发送数据给客户端
func (conn *Connection) SendMsg(msgId uint32, data []byte) error {
	if conn.isClosed {
		return errors.New("Connection closed when send msg")
	}
	// 将data封包,并且发送
	dp := NewDataPack()
	msg, err := dp.Pack(NewMsgPackage(msgId, data))
	if err != nil {
		return errors.New("pack error msg")
	}
	// 写回客户端
	_, err = conn.Conn.Write(msg)
	if err != nil {
		conn.ExitBuffChan <- true
		return errors.New("conn write error")
	}
	return nil
}
