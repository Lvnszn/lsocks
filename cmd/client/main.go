package main

import (
	"fmt"
	"log"
	"lsocks/cmd"
	"lsocks/pkg/client"
	"net"
)

const (
	DefaultListenAddr = ":7448"
)

var version = "master"

func main() {
	log.SetFlags(log.Lshortfile)

	// 默认配置
	config := &cmd.Config{
		ListenAddr: DefaultListenAddr,
	}
	config.ReadConfig()
	config.SaveConfig()

	// 启动 client 端并监听
	lsLocal, err := client.NewLsLocal(config.Password, config.ListenAddr, config.RemoteAddr)
	if err != nil {
		log.Fatalln(err)
	}
	log.Fatalln(lsLocal.Listen(func(listenAddr *net.TCPAddr) {
		log.Println(fmt.Sprintf(`
lightsocks-client:%s 启动成功，配置如下：
本地监听地址：
%s
远程服务地址：
%s
密码：
%s`, version, listenAddr, config.RemoteAddr, config.Password))
	}))
}
