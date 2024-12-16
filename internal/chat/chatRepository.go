package chat

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type chatRepo struct {
	DB *mongo.Collection
}

func NewChatRepo(DB *mongo.Collection) *chatRepo {
	return &chatRepo{DB: DB}
}

type IChatRepo interface {
	NewChat(ctx context.Context, roomId string, chat ChatContent) error
	GetChatInSendToRepo(ctx context.Context, sendTo string) ([]*ChatContent, error)
}

func (c *chatRepo) NewChat(ctx context.Context, roomId string, chat ChatContent) error {
	//  neu phong chua co
	filter := bson.M{
		"roomId": roomId,
	}
	update := bson.M{
		"$push": bson.M{
			"chat": chat,
		},
	}
	opts := options.Update().SetUpsert(true)

	// Thực hiện lệnh update hoặc chèn mới
	_, err := c.DB.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}
	return nil
}
func (c *chatRepo) GetChatInSendToRepo(ctx context.Context, sendTo string) ([]*ChatContent, error) {
	filter := bson.M{"roomId": sendTo}        // Tìm document bằng _id
	projection := bson.M{"chat": 1, "_id": 0} // Chỉ lấy trường 'chat', loại bỏ _id

	var result struct {
		Chat []*ChatContent `bson:"chat"`
	}
	err := c.DB.FindOne(ctx, filter, options.FindOne().SetProjection(projection)).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result.Chat, nil
}
