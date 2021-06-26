package ziface

/*
IRequest 接口
实际是把客户端请求的链接信息和请求信息包装到了Request里
*/
type IRequest interface {
	Getconnection() Iconnetction //获取链接请求信息
	GetData() []byte             // 获取请求消息的数据
}
