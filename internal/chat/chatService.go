package chat

import (
	"context"
	"video-call-project/internal/room"
)

type chatService struct {
	ChatRepo IChatRepo
	//kiem tra xem user co ton tai trong room hay khong
	Room room.IRoomRepo
}

func NewChatService(chatRepo IChatRepo, room room.IRoomRepo) *chatService {
	return &chatService{ChatRepo: chatRepo, Room: room}
}

func (c *chatService) ChatService(ctx context.Context, roomId, uuid string, chat ChatContent) error {
	//  kiem tra xem use co trong room hay khong

	//  luu vao doan chat

	// them vao database
	err := c.ChatRepo.NewChat(ctx, roomId, chat)
	if err != nil {
		return err
	}
	return nil
}

func (c *chatService) GetChatInRoomService(ctx context.Context, roomId string) ([]*ChatContent, error) {
	//  kiem tra xem user co trong room hay khong

	//  lay lich su tin nhan cua user ra
	listChat, err := c.ChatRepo.GetChatInSendToRepo(ctx, roomId)
	if err != nil {
		return nil, err
	}
	return listChat, nil
}
