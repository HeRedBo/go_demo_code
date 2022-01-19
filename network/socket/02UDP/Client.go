package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// localhost:8889
	conn, err := net.Dial("udp", "127.0.0.1:8889")
	ClientHandleError(err, "net.Dial")

	reader := bufio.NewReader(os.Stdin)
	buffer := make([]byte, 1024)

	for {
		lineBytes , _,  _  := reader.ReadLine()
		conn.Write(lineBytes)
		n, _ := conn.Read(buffer)
		serverMsg := string(buffer[:n])
		if serverMsg == "fuckoff" {
			conn.Write([]byte("呜呼狡兔死走狗烹飞鸟尽良弓藏，吾去也！"))
			break
		}
		fmt.Println("GAME OVER!")
	}

}


func ClientHandleError(err error, when string) {
	if err != nil {
		fmt.Println(err, when)
		os.Exit(1)
	}
}