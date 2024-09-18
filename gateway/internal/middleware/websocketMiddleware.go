package middleware

import (
	"fmt"
	"github.com/Sion-L/devops/core"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type WebSocketMiddleware struct {
	ReadBufferSize  int
	WriteBufferSize int
}

func NewWebSocketMiddleware(readSize, writeSize int) *WebSocketMiddleware {
	return &WebSocketMiddleware{
		ReadBufferSize:  readSize,
		WriteBufferSize: writeSize,
	}
}

func (m *WebSocketMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	hub := core.NewHub()
	go hub.Run()
	return func(w http.ResponseWriter, r *http.Request) {
		// 检查是否为WebSocket请求
		if websocket.IsWebSocketUpgrade(r) {
			conn, err := websocket.Upgrade(w, r, nil, m.ReadBufferSize, m.WriteBufferSize)
			// Maximum message size allowed from peer.
			// conn.SetReadLimit(512)
			// other settings..
			// conn.SetReadDeadline(time.Now().Add(core.PongWait))
			// conn.SetPongHandler(func(string) error { conn.SetReadDeadline(time.Now().Add(core.PongWait)); return nil })
			if err != nil {
				httpx.ErrorCtx(r.Context(), w, fmt.Errorf("无法升级到WebSocket: %v", err))
				return
			}
			client := &core.Client{Conn: conn, Send: make(chan []byte, 256)}
			hub.Register <- client

			go m.writePump(client)
			go m.readPump(client, hub, w)

			next(w, r)
		}
	}
}

func (m *WebSocketMiddleware) readPump(c *core.Client, hub *core.Hub, w http.ResponseWriter) {
	defer func() {
		hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				// 记录错误日志
				httpx.Error(w, fmt.Errorf("websocket读取消息失败: %v", err))
			}
			break
		}
		hub.Broadcast <- message
	}
}

func (m *WebSocketMiddleware) writePump(c *core.Client) {
	defer c.Conn.Close()

	for {
		select {
		// 此处是将从客户端接收的消息直接返回,可根据不同的接口写入不同的消息,根据需求改
		case message, ok := <-c.Send:
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			if err := w.Close(); err != nil {
				return
			}
		}
	}
}
