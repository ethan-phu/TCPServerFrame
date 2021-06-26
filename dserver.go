package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"zinx/znet"
)

// 只是负责测试datapack拆包，封包功能
func main() {
	// 创建socket TCP server
	listener, err := net.Listen("tcp", "127.0.0.1:7777")
	if err != nil {
		log.Fatal(err)
		return
	}
	// 创建服务器goroutine,负责从客户端goroutine读取粘包的数据，然后进行解析
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// 处理客户端的请求
		go func(con net.Conn) {
			// 创建拆包封包对象dp
			dp := znet.NewDataPack()
			for {
				// 1. 先读出流中的head部分
				headData := make([]byte, dp.GetHeadLen())
				_, err := io.ReadFull(conn, headData)
				if err != nil {
					log.Fatal(err)
					break
				}
				// 将headData字节流拆包到msg中
				msgHead, err := dp.Unpack(headData)
				if err != nil {
					log.Fatal(err)
					return
				}
				if msgHead.GetDataLen() > 0 {
					// msg是有data数据的，需要再次读取data数据
					// msg := msgHead.(*znet.Message) //??为何这么操作
					msg := &znet.Message{
						Id:      msgHead.GetMsgId(),
						DataLen: msgHead.GetDataLen(),
					}
					msg.Data = make([]byte, msg.DataLen)
					// 根据datalen从io中读取字节流
					_, err := io.ReadFull(conn, msg.Data)
					if err != nil {
						log.Fatal(err)
						return
					}
					fmt.Println("==> Recv Msg:ID=", msg.Id, ",len=", msg.DataLen, "data=", string(msg.Data))
				}
			}
		}(conn)
	}
}
