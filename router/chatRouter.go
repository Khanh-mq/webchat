package router

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	chat2 "video-call-project/internal/chat"
	"video-call-project/internal/room"
	"video-call-project/pkg/middleware"
)

func ChatRouter(r *gin.Engine, db *mongo.Collection, dbRoom *mongo.Collection) {
	//

	newRoomRepo := room.NewRoomRepo(dbRoom)
	newChatRepo := chat2.NewChatRepo(db)
	newSer := chat2.NewChatService(newChatRepo, newRoomRepo)
	newHand := chat2.NewChatHandler(*newSer)
	// gab middleware cung cap mot so thong tin
	chat := r.Group("/chat", middleware.AuthMiddleware(), middleware.CheckUserInRoomMiddleware(newRoomRepo))
	{
		chat.GET("/:roomId", newHand.ConnectHandler)
		go newHand.HandleMessages()
	}
	message := r.Group("/message", middleware.AuthMiddleware(), middleware.CheckUserInRoomMiddleware(newRoomRepo))
	{
		message.GET("/:roomId", newHand.GetChatHand)
	}
}
