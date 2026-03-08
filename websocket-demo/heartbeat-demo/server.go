package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// -------------------------- 常量定义 --------------------------
const (
	// 心跳间隔
	clientPingPeriod = 25 * time.Second // 客户端心跳间隔
	serverPingPeriod = 30 * time.Second // 服务端兜底心跳间隔
	pongWait         = 60 * time.Second // 心跳超时时间
)

// -------------------------- 数据结构 --------------------------
// PingMsg 客户端模拟Ping的消息结构
type PingMsg struct {
	Type string `json:"type"` // "ping"
	Data string `json:"data"` // 心跳数据
}

// PongMsg 服务端回复Pong的消息结构
type PongMsg struct {
	Type string `json:"type"` // "pong"
	Data string `json:"data"` // 心跳数据
}

// Client 单个WebSocket客户端
type Client struct {
	conn    *websocket.Conn // 连接实例
	userId  string          // 用户ID
	mu      sync.Mutex      // 写操作锁（防止并发写连接）
	isClose bool            // 连接是否已关闭
}

// -------------------------- 连接升级器 --------------------------
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 开发环境，生产需限制跨域
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// -------------------------- Client 核心方法 --------------------------
// safeWriteMessage 安全写消息（加锁 + 连接状态检查）
func (c *Client) safeWriteMessage(msgType int, data []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	// 检查连接是否已关闭
	if c.isClose {
		return websocket.ErrCloseSent
	}

	// 设置写超时
	c.conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	return c.conn.WriteMessage(msgType, data)
}

// close 优雅关闭连接
func (c *Client) close() {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.isClose {
		return
	}

	c.isClose = true
	// 发送关闭帧
	_ = c.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	_ = c.conn.Close()

	log.Printf("用户[%s]连接已优雅关闭", c.userId)
}

// serverHeartbeat 服务端主动心跳（兜底）
func (c *Client) serverHeartbeat() {
	ticker := time.NewTicker(serverPingPeriod)
	defer ticker.Stop()

	for range ticker.C {
		// 发送原生Ping帧（浏览器自动响应Pong，不会暴露到onmessage）
		err := c.safeWriteMessage(websocket.PingMessage, []byte(`server_ping_`+time.Now().Format("15:04:05")))
		if err != nil {
			log.Printf("用户[%s]服务端心跳发送失败: %v", c.userId, err)
			c.close()
			return
		}
		log.Printf("用户[%s]服务端发送原生Ping心跳", c.userId)
	}
}

// handlePong 处理Pong响应（无论是客户端还是服务端Ping的响应）
func (c *Client) handlePong() {
	// 设置Pong回调（重置读超时）
	c.conn.SetPongHandler(func(appData string) error {
		log.Printf("用户[%s]收到Pong响应，数据: %s，重置读超时", c.userId, appData)
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})
}

// readPump 读取客户端消息（包括模拟Ping）
func (c *Client) readPump() {
	defer c.close()

	// 初始化读超时
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	// 处理Pong响应
	c.handlePong()

	for {
		msgType, msgBytes, err := c.conn.ReadMessage()
		if err != nil {
			// 区分正常关闭和异常错误
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("用户[%s]读取消息异常: %v", c.userId, err)
			} else {
				log.Printf("用户[%s]连接正常关闭: %v", c.userId, err)
			}
			break
		}

		// 1. 处理原生Ping消息（极少情况，客户端一般用模拟Ping）
		if msgType == websocket.PingMessage {
			log.Printf("用户[%s]收到原生Ping，回复Pong", c.userId)
			_ = c.safeWriteMessage(websocket.PongMessage, msgBytes)
			continue
		}

		// 2. 处理文本消息（业务消息 + 模拟Ping）
		if msgType == websocket.TextMessage {
			// 尝试解析为模拟Ping消息
			var pingMsg PingMsg
			if err := json.Unmarshal(msgBytes, &pingMsg); err == nil && pingMsg.Type == "ping" {
				// 客户端模拟Ping，回复模拟Pong
				log.Printf("用户[%s]收到客户端模拟Ping: %s", c.userId, pingMsg.Data)
				pongMsg := PongMsg{
					Type: "pong",
					Data: pingMsg.Data,
				}
				pongBytes, _ := json.Marshal(pongMsg)
				_ = c.safeWriteMessage(websocket.TextMessage, pongBytes)
				continue
			}

			// 普通业务消息
			log.Printf("用户[%s]收到业务消息: %s", c.userId, string(msgBytes))
			// 回复echo消息
			reply := []byte("服务端回复: " + string(msgBytes))
			_ = c.safeWriteMessage(websocket.TextMessage, reply)
		}
	}
}

// -------------------------- HTTP处理器 --------------------------
func serveWs(w http.ResponseWriter, r *http.Request) {
	// 获取用户ID
	userId := r.URL.Query().Get("userId")
	if userId == "" {
		http.Error(w, "userId不能为空", http.StatusBadRequest)
		return
	}

	// 升级连接
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("升级连接失败: %v", err)
		return
	}

	// 创建客户端
	client := &Client{
		conn:   conn,
		userId: userId,
	}

	log.Printf("用户[%s]连接成功", userId)

	// 启动服务端心跳协程
	go client.serverHeartbeat()
	// 启动读消息协程
	go client.readPump()
}

// -------------------------- 主函数 --------------------------
func main() {
	// 注册路由
	http.HandleFunc("/ws", serveWs)

	// 启动服务
	log.Println("服务端启动: http://localhost:8090")
	log.Fatal(http.ListenAndServe(":8090", nil))
}
