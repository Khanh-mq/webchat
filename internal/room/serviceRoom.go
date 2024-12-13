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
	//  kiem tra xem co ton tai hay khong
	if role != user.Admin {
		roles, exists, _ := r.roomRepo.CheckExistsRoomRepo(c, roomId, uuidUser)
		if exists == false {
			return errors.New("user don't exists in room")
		}
		log.Printf("role user in room : %v  ", roles.Role)
		if roles.Role == user.Admin {
			_, exists, err := r.roomRepo.GetRoomByIdRepo(c, roomId)
			if exists == false {
				return err
			}
			err = r.roomRepo.UpdateRoomRepo(c, roomId, update)
			if err != nil {
				return err
			}
		} else {
			return errors.New("role not exists in room")
		}
		return nil
	} else {
		return errors.New("role not exists in room")
	}

}
