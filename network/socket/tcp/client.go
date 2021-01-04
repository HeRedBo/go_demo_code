package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

/*
·拨号连接服务端主机的8888端口，建立连接
·循环从标准输入（命令行）读取一行用户输入，向服务端发送
·接收并打印服务端消息，如果消息是“bye”，就退出程序
*/

func ClientHandleError(err error, when string) {
	if err != nil {
		fmt.Println(err, when)
		os.Exit(1)
	}
}

func main() {

	conn, err := net.Dial("tcp","127.0.0.1:8888")
	ClientHandleError(err, "net.Dial")

	buffer := make([]byte, 1024)

	reader := bufio.NewReader(os.Stdin)

	//源源不断地收发消息
	for {
		/*一个完整的消息回合*/

		//接收命令行输入（的一行消息）
		lineBytes, _, _ := reader.ReadLine()

		//向服务端发送
		conn.Write(lineBytes)

		//接收服务端消息，存入消息缓冲区，长度是n字节
		n, _ := conn.Read(buffer)

		//转化为字符串并打印
		serverMsg := string(buffer[0:n])
		fmt.Println("服务端:", serverMsg)

		//如果服务端说bye就退出消息循环
		if serverMsg == "bye" {
			break
		}
	}

	//客户端程序结束
	fmt.Println("GAME OVER!")


}