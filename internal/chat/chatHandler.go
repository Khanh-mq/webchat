package chat

import (
	"errors"
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
func (c *chatHandler) sendHistoryChatHandler(ctx *gin.Context, conn *websocket.Conn) error {
	roomId := ctx.Param("roomId")
	// lay lich su tin nhan
	history, err := c.chat.ChatRepo.GetChatInSendToRepo(ctx, roomId)
	if err != nil {
		return err
	}
	//  gui tin nhan vao cho cac memeber in room
	err = conn.WriteJSON(history)
	if err != nil {
		return err
	}
	return nil
}
func (c *chatHandler) registerClient(ctx *gin.Context, conn *websocket.Conn) error {
	//  kiem tra xem use co trong database hay chua (do user exists in room ? )
	uuid := ctx.MustGet("uuid").(string)
	roomId := ctx.Param("roomId")
	if _, exists, _ := c.chat.Room.CheckExistsRoomRepo(ctx, roomId, uuid); exists == false {
		return errors.New("use dont exists in room")
	}
	if _, exists := clients[conn]; exists {
		log.Println("client is already registered")
		return errors.New("client is already registered")
	}
	clients[conn] = true
	log.Println("client is registered")
	return nil
}
func (c *chatHandler) unregisterClient(conn *websocket.Conn) error {
	//  thoat ket loi  cua user
	delete(clients, conn)
	conn.Close()
	return nil
}
func (c *chatHandler) inComingMessageHandler(ctx *gin.Context, conn *websocket.Conn) error {
	//  gui tin nhan den cho database de luu voa trong csdl

	roomId := ctx.Param("roomId")
	uuid := ctx.MustGet("uuid").(string)
	//  doc tin nhan tu userMessage
	for {
		var text ChatText
		err := conn.ReadJSON(&text)
		if err != nil {
			return err
		}
		msg := ChatContent{
			Content:  text.Content,
			UserName: uuid,
			TimeChat: time.Now(),
		}
		//  luu vao csdl cho user
		err = c.chat.ChatRepo.NewChat(ctx, roomId, msg)
		if err != nil {
			return err
		}
		//  chuyen tin nhan ve cho broadcast
		broadcast <- msg
	}
}
func (c *chatHandler) broadcastSendMessageToClient(msg ChatContent) {
	for conn := range clients {
		err := conn.WriteJSON(msg)
		if err != nil {
			log.Println(err)
			return
		}
	}
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
	//  kiem tra ton tai va cho vao trong phong
	err = c.registerClient(ctx, conn)
	if err != nil {
		log.Println(err)
		return
	}
	// Close connection when function ends
	defer func() {
		delete(clients, conn)
		conn.Close()
	}()
	// sau khi ket noi , gui tin nhan cu cho user khi connect room
	// ty thưcj hiện kiểm tra quyền trong middleware

	// lay lich su tin nhan ra
	err = c.sendHistoryChatHandler(ctx, conn)
	if err != nil {
		log.Println(err)
		return
	}
	//  su ly  nhan tin va luu vao csdl
	err = c.inComingMessageHandler(ctx, conn)
	if err != nil {
		log.Println(err)
		return
	}

}
func (c *chatHandler) HandleMessages() {
	for {
		// Nhận tin nhắn từ kênh broadcast
		msg := <-broadcast
		// Gửi tin nhắn đến tất cả client đã kết nối
		c.broadcastSendMessageToClient(msg)
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
