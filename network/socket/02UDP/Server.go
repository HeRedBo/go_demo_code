package main

import (
	"fmt"
	"net"
	"os"
)

/*
·服务端在本机的8888端口建立UDP监听，得到广口连接
·循环接收客户端消息，不管客户端说什么，都自动回复“已阅xxx”
·如果客户端说的是“im off”，则回复“bye”
*/
func main() {
	// 解析得到UDP地址
	udpAddr , err := net.ResolveUDPAddr("udp","localhost:8889")
	ServerHandleError(err, "net.ResolveUDPAddr")

	// 建立 UDP 监听，得到广口连接
	udpConn, err := net.ListenUDP("udp", udpAddr)
	ServerHandleError(err, "net.ListenUDP")

	// 创建爱消息缓冲区
	buffer := make([]byte,1024)
	// 从广口连接中源源不断读取 （来自任何客户端的）数据包
	for {
		n , remoteAddr , _ := udpConn.ReadFromUDP(buffer)
		clientMsg := string(buffer[:n])
		fmt.Println("收到来自%v的消息：%s \n", remoteAddr, clientMsg)

		if clientMsg != "im off" {
			 udpConn.WriteToUDP([]byte("已阅：" + clientMsg), remoteAddr)
		} else {
			udpConn.WriteToUDP([]byte("fuckoff"), remoteAddr)
		}
	}
}

func ServerHandleError(err error, when string) {
	if err != nil {
		fmt.Println(err, when)
		os.Exit(1)
	}
}
