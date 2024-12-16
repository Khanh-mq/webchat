package chat

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

// nang cap webchat
var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true // Allow all origins (adjust for production)
		},
	}
	clients   = make(map[*websocket.Conn]bool)
	broadcast = make(chan ChatContent)
)

type chatHandler struct {
	chat chatService
}

func NewChatHandler(chat chatService) *chatHandler {
	return &chatHandler{chat: chat}
}

// sau khi chuyen  cai dat duoc vao thi ket hop voi csdl
func (c *chatHandler) ConnectHandler(ctx *gin.Context) {
	// Upgrade HTTP request to WebSocket
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Printf("WebSocket upgrade error: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upgrade connection to WebSocket"})
		return
	}

	// Check if user is already in a room
	if _, exists := clients[conn]; exists {
		log.Printf("Client already connected: %v", conn.RemoteAddr())
		conn.Close()
		return
	}

	// Register new client
	clients[conn] = true

	// Close connection when function ends
	defer func() {
		delete(clients, conn)
		conn.Close()
	}()
	// sau khi ket noi , gui tin nhan cu cho user khi connect room
	// ty thưcj hiện kiểm tra quyền trong middleware
	roomId := ctx.Param("roomId")
	uuid, exists := ctx.MustGet("uuid").(string)
	if !exists {
		log.Printf("UUID not found in context")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "UUID not found in context"})
		return
	}

	historyMessage, err := c.chat.GetChatInRoomService(ctx, roomId)
	if err != nil {
		log.Printf("Failed to get history message: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get history message"})
		return
	}

	// Read messages from this connection
	// hien tin nhan cu cho user
	err = conn.WriteJSON(historyMessage)
	if err != nil {
		log.Printf("Failed to send history message: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send history message"})
		return
	}
	for {
		var text ChatText

		err := conn.ReadJSON(&text)
		if err != nil {
			log.Printf("Connection read error: %v", err)
			delete(clients, conn)
			break
		}
		var msg = ChatContent{
			Content:  text.Content,
			UserName: uuid,
			TimeChat: time.Now(),
		}
		//  luyw tin nhan vao cơ sở dư liệu
		err = c.chat.ChatService(ctx, roomId, uuid, msg)
		if err != nil {
			log.Printf("Failed to chat message: %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to chat message"})
			return
		}
		// Send received message to broadcast channel
		broadcast <- msg
	}
}
func (c *chatHandler) HandleMessages() {
	for {
		// Nhận tin nhắn từ kênh broadcast
		msg := <-broadcast

		// Gửi tin nhắn đến tất cả client đã kết nối
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("Lỗi khi gửi tin nhắn: %v", err)
				client.Close()
				delete(clients, client) // Xóa client nếu xảy ra lỗi
			}
		}
	}

}
func (c *chatHandler) GetChatHand(ctx *gin.Context) {
	roomId := ctx.Param("roomId")
	listChat, err := c.chat.GetChatInRoomService(ctx, roomId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get history message"})
		return

	}
	ctx.JSON(http.StatusOK, listChat)

}
