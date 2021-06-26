package znet

import "zinx/ziface"

type Message struct {
	Id      uint32 // 消息的ID
	DataLen uint32 //消息的长度
	Data    []byte // 消息内容
}

func NewMsgPackage(id uint32, data []byte) ziface.IMessage {
	return &Message{
		Id:      id,
		DataLen: uint32(len(data)),
		Data:    data,
	}
}

// 获取消息数据长度
func (msg *Message) GetDataLen() uint32 {
	return msg.DataLen
}

// 获取消息ID
func (msg *Message) GetMsgId() uint32 {
	return msg.Id
}

// 获取消息
func (msg *Message) GetData() []byte {
	return msg.Data
}

// 设计消息ID
func (msg *Message) SetMsgId(id uint32) {
	msg.Id = id
}

//设计消息内容
func (msg *Message) SetData(data []byte) {
	msg.Data = data
}

// 设置消息数据段长度
func (msg *Message) SetDataLen(len uint32) {
	msg.DataLen = len
}
