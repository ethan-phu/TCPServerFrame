package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"zinx/utils"
	"zinx/ziface"
)

// 封包拆包类实例
type DataPack struct {
}

// 封包拆包实例初始化方法
func NewDataPack() ziface.IDataPack {
	return &DataPack{}
}

// 获取包头方法
func (dp *DataPack) GetHeadLen() uint32 {
	// Id uint32(4字节) + DataLen uint32(4字节)
	return 8
}

// 封包数据
func (dp *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {
	// 创建一个存放bytes字节的缓冲
	dataBuff := bytes.NewBuffer([]byte{})
	// 写dataLen
	err := binary.Write(dataBuff, binary.LittleEndian, msg.GetDataLen())
	if err != nil {
		return nil, err
	}
	// 写msgId
	err = binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgId())
	if err != nil {
		return nil, err
	}
	// 写数据
	err = binary.Write(dataBuff, binary.LittleEndian, msg.GetData())
	if err != nil {
		return nil, err
	}
	return dataBuff.Bytes(), nil
}

// 拆包方法
func (dp *DataPack) Unpack(binaryData []byte) (ziface.IMessage, error) {
	// 创建一个DataBuff
	dataBuff := bytes.NewReader(binaryData)
	// 只解压head的信息，得到dataLen和msgID
	msg := &Message{}
	// 解压ID
	err := binary.Read(dataBuff, binary.LittleEndian, &msg.DataLen)
	if err != nil {
		return nil, err
	}
	// 解压ID
	err = binary.Read(dataBuff, binary.LittleEndian, &msg.Id)
	if err != nil {
		return nil, err
	}
	// 判断dataLen的长度是否超出我们允许的最大包长度
	if utils.GlobalObject.MaxPacketSize > 0 && msg.DataLen > utils.GlobalObject.MaxPacketSize {
		return nil, errors.New("Too large msg data reciedved")
	}
	// 这里只需把head的数据拆包出来就可以了，然后在通过head的长度，再从conn读取一次数据
	return msg, nil
}
