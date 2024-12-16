package room

import (
	"context"
	"errors"
	"log"
	"video-call-project/internal/user"
)

type roomService struct {
	roomRepo IRoomRepo
}

func NewRoomService(roomRepo IRoomRepo) *roomService {
	return &roomService{roomRepo: roomRepo}
}
func (r *roomService) CreateRoomSer(c context.Context, newRoom RoomCreate, userCreate string) error {
	//  su ly  room trong nay

	rooms := NewRoom(newRoom.RoomName, newRoom.TotalMember, userCreate)
	err := r.roomRepo.CreateRoomRepo(c, *rooms)
	if err != nil {
		return err
	}
	return nil
}
func (r *roomService) GetListRoomSer(c context.Context) ([]*rooms, error) {
	// lay danh sach
	result, err := r.roomRepo.GetListRoomRepo(c)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (r *roomService) joinRoomSer(c context.Context, roomName, uuid string) error {
	// kiem tra xem user co ton tai trong room hay chua
	_, exists, _ := r.roomRepo.CheckExistsRoomRepo(c, roomName, uuid)
	if exists == true {
		return errors.New("user exists in room")
	}
	newMember := NewMember(uuid)
	//  cho vao room
	err := r.roomRepo.JoinRoomRepo(c, roomName, *newMember)
	if err != nil {
		return errors.New("join room fail")
	}
	return nil
}
func (r *roomService) GetUserInRoomSer(c context.Context, roomId string, userId string, role interface{}) ([]MemberRole, error) {
	//  kiem tra xem user co  trong room do hay khong
	// néu là admin-system thi khong can kiem tra
	//  phai kiem tra 2  phan
	if role != user.Admin {
		_, exists, _ := r.roomRepo.CheckExistsRoomRepo(c, roomId, userId)
		if exists == false {
			return nil, errors.New("user don't exists in room")
		}
	}
	//  lay danh sach ra
	result, err := r.roomRepo.GetUserInRoomRepo(c, roomId)
	if err != nil {
		return nil, err
	}
	return result, nil

}
func (r *roomService) UpdateRoomSer(c context.Context, roomId string, update updateRoom, role interface{}, uuidUser string) error {
	_, exists, _ := r.roomRepo.GetRoomByIdRepo(c, roomId)
	if exists == false {
		return errors.New("user don't exists in room")
	}
	if err := r.roomRepo.UpdateRoomRepo(c, roomId, update); err != nil {
		return err
	}
	return nil
}

func (r *roomService) DeleteRoomSer(c context.Context, roomId string) error {
	// o day ta len thuc hien khi xoa phong thi nen xoa luon tin nhan cua phong  di nua
	// kiem tra xem phong co ton tai hay khong
	_, exists, err := r.roomRepo.GetRoomByIdRepo(c, roomId)
	if err != nil || exists == false {
		return errors.New("room don't exists ")

	}
	//  xoa phong do
	err = r.roomRepo.DeletedRoomRepo(c, roomId)
	if err != nil {
		return err
	}
	return nil
}
func (r *roomService) AddUserInRoomSer(c context.Context, roomId string, user1 MemberRole) error {
	//  kiem tra xem user co trong room hay chua
	_, exists, _ := r.roomRepo.CheckExistsRoomRepo(c, roomId, user1.UserId)
	if exists == true {
		// neu ton tai
		return errors.New("user exists in room")
	}
	user1.Role = user.Member
	//  chua ton tai thi them vao trong room
	err := r.roomRepo.AddUserInRoomRepo(c, roomId, user1)
	if err != nil {
		return err
	}
	return nil
}
func (r *roomService) DeletedUserInRoomSer(c context.Context, roomId string, uuid string) error {
	err := r.roomRepo.DeletedUserInRoom(c, roomId, uuid)
	if err != nil {
		return err
	}
	return nil
}
func (r *roomService) UpdateRoleInRoomSer(c context.Context, roomId, uuid string, role string) error {
	// chuyen role sang item role
	// kiem tra xem user co trong room hay khong
	_, exists, _ := r.roomRepo.CheckExistsRoomRepo(c, roomId, uuid)
	if exists == false {
		return errors.New("user dont exists in room ")
	}
	roleItem, err := user.ParseStr2ItemRole(role)
	if err != nil {
		return err
	}
	log.Println(roleItem)
	err = r.roomRepo.UpdateRoleRepo(c, roomId, uuid, roleItem)
	if err != nil {
		return err
	}
	return nil
}
