package ziface

/*
消息管理层
*/

type IMessageHandle interface {
	DoMsgHandle(request IRequest)
	AddRouter(msgId uint32, router IRouter)
}
