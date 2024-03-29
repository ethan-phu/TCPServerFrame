package ziface

type Iserver interface {
	// 启动服务器
	Start()
	// 停止服务器
	Stop()
	// 运行服务器
	Serve()
	// 路由功能：给当前服务注册一个路由业务方法，给客户端处理使用
	AddRouter(msgId uint32, router IRouter)
}
