package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 创建一个缓冲大小为 1 的 os.Signal 类型的通道
	sigterm := make(chan os.Signal, 1)
	// 监听 SIGINT 和 SIGTERM 信号  SIGINT（中断）和 SIGTERM（终止）
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("程序开始运行，按 Ctrl+C 或者发送 SIGTERM 信号来终止程序...")

	// 阻塞等待信号
	sig := <-sigterm
	fmt.Printf("收到信号: %v，程序即将终止...\n", sig)
}
