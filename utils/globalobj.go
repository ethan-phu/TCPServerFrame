package utils

import (
	"encoding/json"
	"io/ioutil"
	"zinx/ziface"
)

/*
存储一切有关Zinx框架的全局参数，供其他模块使用
一些参数也可以通过用户根据zinx.json来配置
*/
type GlobalObj struct {
	TcpServer ziface.Iserver //当前Zinx的全局server对象
	Host      string         // 当前服务器主机的IP
	TcpPort   int            // 当前服务器主机的端口
	Name      string         // 当前服务器的名称
	Version   string         // 当前Zinx版本号

	MaxPacketSize uint32 // 数据包的最大值
	MaxConn       int    // 当前服务器主机允许的最大连接数
}

var GlobalObject *GlobalObj

// 读取用户的配置文件
func (g *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("conf/zinx.json")
	if err != nil {
		panic(err)
	}
	// 将json数据解析到struct中
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}

// 提供一个Init方法，默认加载
func init() {
	GlobalObject = &GlobalObj{
		Name:          "ZinxServerApp",
		Version:       "V0.4",
		TcpPort:       8889,
		Host:          "0.0.0.0",
		MaxConn:       12000,
		MaxPacketSize: 4096,
	}
	// 从配置文件中加载一些用户配置的参数
	GlobalObject.Reload()
}
