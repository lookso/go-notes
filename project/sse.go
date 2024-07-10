package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func sseHandler(w http.ResponseWriter, r *http.Request) {
	// 设置响应头为 Server-Sent Events 标准格式
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	ctx := r.Context() // 获取请求的上下文
	// 模拟数据流，定时发送消息
	for i := 0; ; i++ {
		select {
		case <-ctx.Done(): // 检查上下文是否被取消
			log.Println("Client disconnected")
			return
		default:
			// 构造消息
			msg := fmt.Sprintf("data: Message %d\n\n", i)

			// 发送消息
			_, err := w.Write([]byte(msg))
			if err != nil {
				// 如果无法发送消息，比如客户端断开连接，就结束循环
				log.Printf("Error writing to client: %v", err)
				break
			}

			// 刷新缓冲区，确保消息立即发送
			if f, ok := w.(http.Flusher); ok {
				f.Flush()
			}

			// 等待一段时间再发送下一条消息
			time.Sleep(2 * time.Second)
		}
	}
}

func main() {
	http.HandleFunc("/events", sseHandler)

	log.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
