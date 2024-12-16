package chat

import "time"

// su ly nếu là nhóm thì phải gom vào một tin nhắn  chỗ
// tương tự với user cũng thế

type ChatContent struct {
	Content  string    `json:"content" bson:"content"`
	UserName string    `json:"userName" bson:"userName"`
	TimeChat time.Time `json:"timeChat" bson:"timeChat"`
}

func NewChatContent(content string, userName string) *ChatContent {
	return &ChatContent{Content: content, UserName: userName, TimeChat: time.Now()}
}

type ChatRoom struct {
	Chat []*ChatContent `json:"chat" bson:"chat"`
}

func NewChat(chat []*ChatContent) *ChatRoom {
	return &ChatRoom{Chat: chat}
}

type ChatText struct {
	Content string `json:"content" bson:"content"`
}
