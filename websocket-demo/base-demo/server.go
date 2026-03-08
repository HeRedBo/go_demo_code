package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

// 定义升级器，用于将HTTP连接升级为WebSocket连接
var upgrader = websocket.Upgrader{
	// 允许跨域（开发环境临时开启，生产环境需限制）
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// 心跳检测间隔（30秒）
const (
	pingPeriod = 30 * time.Second
)

// 处理WebSocket连接的核心函数
func wsHandler(w http.ResponseWriter, r *http.Request) {
	// 1. 将HTTP连接升级为WebSocket连接
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("升级连接失败: %v", err)
		return
	}
	// 确保连接关闭
	defer func() {
		log.Println("连接关闭")
		conn.Close()
	}()

	// 2. 设置连接的读写超时
	conn.SetReadDeadline(time.Now().Add(pingPeriod * 2))
	// 3. 设置Pong回调（客户端响应心跳后重置超时）
	conn.SetPongHandler(func(string) error {
		conn.SetReadDeadline(time.Now().Add(pingPeriod * 2))
		return nil
	})

	// 4. 启动心跳协程（服务端主动发Ping）
	go func() {
		ticker := time.NewTicker(pingPeriod)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				// 发送Ping帧，检测客户端是否在线
				if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
					log.Printf("发送心跳失败: %v", err)
					return
				}
			}
		}
	}()

	// 5. 循环读取客户端消息
	for {
		// 读取消息（类型：文本/二进制/Ping/Pong/Close）
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Printf("读取消息失败: %v", err)
			break
		}

		// 打印客户端消息
		log.Printf("收到客户端消息: %s (类型: %d)", msg, msgType)

		// 6. 回复客户端（echo功能）
		if err := conn.WriteMessage(msgType, msg); err != nil {
			log.Printf("回复消息失败: %v", err)
			break
		}
	}
}

func main() {
	// 注册WebSocket路由
	http.HandleFunc("/ws", wsHandler)
	// 启动HTTP服务
	log.Println("服务端启动: http://localhost:8088")
	log.Fatal(http.ListenAndServe(":8088", nil))
}
