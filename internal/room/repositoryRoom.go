package room

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type roomRepo struct {
	DB *mongo.Collection
}

func NewRoomRepo(db *mongo.Collection) *roomRepo {
	return &roomRepo{
		DB: db,
	}
}

//  ham kiem tra su ton tai cua room hay chua
//  xem  cos ten nao trung voi ten room  muon tao  hay khong?

func (r *roomRepo) GetRoomByNameRepo(c context.Context, roomName string) (bool, error) {
	err := r.DB.FindOne(c, bson.M{"roomName": roomName}).Err()
	if errors.Is(err, mongo.ErrNoDocuments) {
		//  khong tim thay tra ve false
		return false, err
	}
	return true, nil
}

// truong hop nay  neu ma
func (r *roomRepo) CreateRoomRepo(c context.Context, newRoom rooms) error {
	// check exists
	exists, _ := r.GetRoomByNameRepo(c, newRoom.RoomName)
	if exists == true {
		return errors.New("room was exists")
	}
	_, err := r.DB.InsertOne(c, newRoom)
	if err != nil {
		return err
	}
	return nil
}
func (r *roomRepo) GetListRoomRepo(c context.Context) ([]*ViewRoom, error) {
	cursor, err := r.DB.Find(c, bson.M{})
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {

		}
	}(cursor, c)
	var result []*ViewRoom
	err = cursor.All(c, &result)
	if err != nil {
		return nil, err
	}
	return result, nil

}
func (r *roomRepo) CheckExistsRoomRepo(c context.Context, roomId, uuid string) (bool, error) {
	filter := bson.M{"roomId": roomId,
		"roleMember": bson.M{
			"$elemMatch": bson.M{
				"userId": uuid,
			},
		},
	}
	err := r.DB.FindOne(c, filter).Err()
	if errors.Is(err, mongo.ErrNoDocuments) {
		// neu khong tin tai thi
		return false, errors.New("user does not exist in room ")
	}
	return true, nil

}
func (r *roomRepo) JoinRoomRepo(c context.Context, roomName string, member MemberRole) error {
	filter := bson.M{"roomId": roomName}
	update := bson.M{"$push": bson.M{"roleMember": member}}
	_, err := r.DB.UpdateOne(c, filter, update)
	if err != nil {
		return errors.New("failed to join room")
	}
	return nil
}
func (r *roomRepo) GetUserInRoomRepo(c context.Context, roomId string) ([]MemberRole, error) {
	var results []MemberRole
	filter := bson.M{
		"roomId": roomId,
	}

	cursor, err := r.DB.Find(c, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)

	for cursor.Next(c) {
		var member MemberRole
		err := cursor.Decode(&member)
		if err != nil {
			return nil, err
		}
		results = append(results, member)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	log.Println(results)
	return results, nil
}

type IRoomRepo interface {
	CreateRoomRepo(c context.Context, newRoom rooms) error
	GetListRoomRepo(c context.Context) ([]*ViewRoom, error)
	JoinRoomRepo(c context.Context, roomName string, member MemberRole) error
	CheckExistsRoomRepo(c context.Context, roomId, uuid string) (bool, error)
	GetUserInRoomRepo(c context.Context, roomId string) ([]MemberRole, error)
}
