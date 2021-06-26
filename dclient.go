package main

import (
	"log"
	"net"
	"zinx/znet"
)

func main() {
	//从客户端goroutine，负责模拟粘包的数据，然后进行发送
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		log.Fatal(err)
		return
	}
	// 创建一个封包对象 dp
	dp := znet.NewDataPack()
	// 封包一个msg1包
	msg1 := &znet.Message{
		Id:      0,
		DataLen: 5,
		Data:    []byte{'h', 'e', 'l', 'l', 'o'},
	}
	msgbyte1, err := dp.Pack(msg1)
	if err != nil {
		log.Fatal(err)
		return
	}
	// 封包一个msg1包
	msg2 := &znet.Message{
		Id:      1,
		DataLen: 5,
		Data:    []byte{'w', 'o', 'r', 'l', 'd'},
	}
	msgbyte2, err := dp.Pack(msg2)
	if err != nil {
		log.Fatal(err)
		return
	}
	msgbyte1 = append(msgbyte1, msgbyte2...)
	//向服务端发送数据
	conn.Write(msgbyte1)
	// 客户端阻塞
	select {}
}
