package room

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"video-call-project/internal/user"
)

type roomRepo struct {
	DB *mongo.Collection
}

func NewRoomRepo(db *mongo.Collection) *roomRepo {
	return &roomRepo{
		DB: db,
	}
}
func NewRoomChatRepo(db *mongo.Database) *roomRepo {
	return &roomRepo{
		DB: db.Collection("rooms")}
}

type IRoomRepo interface {
	CreateRoomRepo(c context.Context, newRoom rooms) error
	GetListRoomRepo(c context.Context) ([]*rooms, error)
	JoinRoomRepo(c context.Context, roomName string, member MemberRole) error
	CheckExistsRoomRepo(c context.Context, roomId, uuid string) (*MemberRole, bool, error)
	GetUserInRoomRepo(c context.Context, roomId string) ([]MemberRole, error)
	UpdateRoomRepo(c context.Context, roomId string, updateRoom updateRoom) error
	GetRoomByIdRepo(c context.Context, roomId string) (*ViewRoom, bool, error)
	DeletedRoomRepo(c context.Context, roomId string) error
	AddUserInRoomRepo(c context.Context, roomId string, user MemberRole) error
	DeletedUserInRoom(c context.Context, roomId, uuid string) error
	UpdateRoleRepo(c context.Context, roomId, uuid string, role user.ItemRole) error
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
func (r *roomRepo) DeletedUserInRoom(c context.Context, roomId, uuid string) error {
	filter := bson.M{"roomId": roomId,
		"roleMember.userId": uuid,
	}
	_, err := r.DB.DeleteOne(c, filter)
	if err != nil {
		return err
	}
	return nil
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
func (r *roomRepo) GetListRoomRepo(c context.Context) ([]*rooms, error) {
	cursor, err := r.DB.Find(c, bson.M{})
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {

		}
	}(cursor, c)
	var result []*rooms
	err = cursor.All(c, &result)
	if err != nil {
		return nil, err
	}
	return result, nil

}

// kiem tra user ton tai
func (r *roomRepo) GetRoomByIdRepo(c context.Context, roomId string) (*ViewRoom, bool, error) {
	var view ViewRoom
	err := r.DB.FindOne(c, bson.M{"roomId": roomId}).Decode(&view)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, false, err
	}
	return &view, true, nil
}
func (r *roomRepo) CheckExistsRoomRepo(c context.Context, roomId, uuid string) (*MemberRole, bool, error) {
	// Define filter to find the user within the room's roleMember array
	role, err := r.GetUserInRoomRepo(c, roomId)
	for _, memberRole := range role {
		if memberRole.UserId == uuid {
			return &memberRole, true, nil
		}
	}
	return nil, false, err

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
	reult, err := r.GetListRoomRepo(c)
	if err != nil {
		return nil, err
	}

	for _, res := range reult {
		if res.RoomId == roomId {
			return res.RoleMember, nil
		}
	}
	return nil, errors.New(roomId + " not found in room")

}

func (r *roomRepo) UpdateRoomRepo(c context.Context, roomId string, updateRoom updateRoom) error {
	// user
	data, err := bson.Marshal(updateRoom)
	if err != nil {
		return err
	}
	filter := bson.M{"roomId": roomId}

	updateFiedls := bson.M{}
	err = bson.Unmarshal(data, &updateFiedls)
	if err != nil {
		return err
	}
	//  bo qua nhung truong nil
	for key, value := range updateFiedls {
		if value == nil {
			delete(updateFiedls, key)
		}
	}
	update := bson.M{"$set": updateFiedls}
	_, err = r.DB.UpdateOne(c, filter, update)
	if err != nil {
		return err
	}
	return nil

}
func (r *roomRepo) DeletedRoomRepo(c context.Context, roomId string) error {
	filter := bson.M{"roomId": roomId}
	_, err := r.DB.DeleteOne(c, filter)
	if err != nil {
		return err
	}
	return nil
}
func (r *roomRepo) AddUserInRoomRepo(c context.Context, roomId string, user MemberRole) error {
	filter := bson.M{"roomId": roomId}
	update := bson.M{"$push": bson.M{"roleMember": user}}
	_, err := r.DB.UpdateOne(c, filter, update)
	if err != nil {
		return err
	}
	return nil
}
func (r *roomRepo) UpdateRoleRepo(c context.Context, roomId, uuid string, role user.ItemRole) error {
	filter := bson.M{
		"roomId":     roomId,
		"roleMember": bson.M{"$elemMatch": bson.M{"userId": uuid}},
	}
	// Cập nhật role của user
	update := bson.M{
		"$set": bson.M{"roleMember.$.role": role},
	}
	_, err := r.DB.UpdateOne(c, filter, update)
	if err != nil {
		return err
	}
	return nil
}
