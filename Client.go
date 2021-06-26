package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

/*
模拟客户端
*/
func main() {
	// 1 直接连接远程服务器，得到一个conn连接
	conn, err := net.Dial("tcp4", "127.0.0.1:8999")
	if err != nil {
		log.Fatal(err)
		return
	}
	// 2. 连接调用write数据
	for i := 0; i < 10; i++ {
		_, err := conn.Write([]byte("helle zinx w0.1"))
		if err != nil {
			log.Fatal(err)
			return
		}
		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println("server call back:%s,cnt=%d,content:%s\n", buf, cnt, string(buf[:cnt]))
		time.Sleep(time.Second * 3)

	}
}
