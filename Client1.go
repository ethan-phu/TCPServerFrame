package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
	"zinx/znet"
)

/*
模拟客户端
*/
func main() {
	// 1 直接连接远程服务器，得到一个conn连接
	conn, err := net.Dial("tcp4", "127.0.0.1:8899")
	if err != nil {
		log.Fatal(err)
		return
	}
	// 2. 连接调用write数据
	for i := 0; i < 10; i++ {
		// 发送封包信息
		dp := znet.NewDataPack()
		// 进行封包处理
		msg, err := dp.Pack(znet.NewMsgPackage(1, []byte("hello zinx v0.1")))
		if err != nil {
			log.Fatal(err)
			return
		}
		// 发送信息
		_, err = conn.Write(msg)
		if err != nil {
			log.Fatal(err)
			return
		}
		headData := make([]byte, dp.GetHeadLen())
		_, err = io.ReadFull(conn, headData)
		if err != nil {
			log.Fatal(err)
			return
		}
		msgHead, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("server unpack err:", err)
			return
		}
		if msgHead.GetDataLen() > 0 {
			// msg是有data数据的，需要再次读取data数据
			msg := msgHead.(*znet.Message)
			msg.Data = make([]byte, msg.GetDataLen())
			// 根据dataLen从Io中读取字节流
			_, err := io.ReadFull(conn, msg.Data)
			if err != nil {
				log.Fatal(err)
				return
			}
			fmt.Println("==> Recv Msg: ID=", msg.Id, ", len=", msg.DataLen, ", data=", string(msg.Data))
		}

		time.Sleep(time.Second * 3)

	}
}
