package znet

import "zinx/ziface"

type Request struct {
	conn ziface.Iconnetction //已经和客户端建立好的请求链接
	msg  ziface.IMessage     // 客户端请求的数据
}

// 获取请求连接信息
func (r *Request) Getconnection() ziface.Iconnetction {
	return r.conn
}

// 获取数据
func (r *Request) GetData() []byte {
	return r.msg.GetData()
}

// 获取请求的数据的ID
func (r *Request) GetMsgID() uint32 {
	return r.msg.GetMsgId()
}
