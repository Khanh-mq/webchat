package room

import (
	"github.com/google/uuid"
	"time"
	"video-call-project/internal/user"
)

type rooms struct {
	RoomId      string       `json:"roomId" bson:"roomId"`
	RoomName    string       `json:"roomName" bson:"roomName"`
	TotalMember int          `json:"totalMember" bson:"totalMember"`
	RoleMember  []MemberRole `json:"roleMember" bson:"roleMember"`
	CreateTime  time.Time    `json:"createTime" bson:"createTime"`
}
type MemberRole struct {
	Role   user.ItemRole `json:"role" bson:"role"`
	UserId string        `json:"userId" bson:"userId"`
}

func NewMember(userId string) *MemberRole {
	return &MemberRole{
		UserId: userId,
		Role:   user.Member, //  mac dinh la cho la member in room
	}
}

type ViewRoom struct {
	RoomId      string    `json:"roomId" bson:"roomId"`
	RoomName    string    `json:"roomName" bson:"roomName"`
	TotalMember int       `json:"totalMember" bson:"totalMember"`
	CreateTime  time.Time `json:"createTime" bson:"createTime"`
}

type updateRoom struct {
	RoomName    string    `json:"roomName" bson:"roomName"`
	TotalMember int       `json:"totalMember" bson:"totalMember"`
	CreateTime  time.Time `json:"createTime" bson:"createTime"`
}

func NewUpdateRoom(userRoom string, totalMember int) *updateRoom {
	return &updateRoom{
		RoomName:    userRoom,
		TotalMember: totalMember,
		CreateTime:  time.Now(),
	}

}
func NewRoom(roomName string, totalMember int, userCreate string) *rooms {
	return &rooms{
		RoomId:      uuid.New().String(),
		RoomName:    roomName,
		TotalMember: totalMember,
		RoleMember: []MemberRole{
			MemberRole{
				UserId: userCreate,
				//  people first role is admin
				Role: user.Admin,
			},
		},
		CreateTime: time.Now(),
	}

}

type RoomCreate struct {
	RoomName    string `json:"roomName" bson:"roomName"`
	TotalMember int    `json:"totalMember" bson:"totalMember"`
}
