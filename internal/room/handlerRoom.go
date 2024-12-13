package room

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type roomHandler struct {
	roomSer roomService
}

func NewRoomHandler(roomSer *roomService) *roomHandler {
	return &roomHandler{roomSer: *roomSer}
}
func (r *roomHandler) CreateRoomHand(c *gin.Context) {
	// lay uuid cua usr r
	var newRoom RoomCreate
	err := c.ShouldBind(&newRoom)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": "error json ",
		})
	}
	uuid := c.MustGet("uuid").(string)
	err = r.roomSer.CreateRoomSer(c, newRoom, uuid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": "success",
	})
}
func (r *roomHandler) GetListRoomHand(c *gin.Context) {
	//  oe day cho phep su ly phan trang nua
	// lay phan trang tu json
	data, err := r.roomSer.GetListRoomSer(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
func (r *roomHandler) JoinRoomHand(c *gin.Context) {
	// lay uuid tu token
	uuid := c.MustGet("uuid").(string)
	//  lay roomID tu param
	roommId := c.Param("roomId")

	err := r.roomSer.joinRoomSer(c, roommId, uuid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": "data error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": "success",
	})
	// di chuyen nguoi dung sang  doan chat message
}
func (r *roomHandler) GetUserRoomHand(c *gin.Context) {
	//  lay role in token
	role := c.MustGet("role").(string)
	log.Println(role)
	//  if role la admin thi   truy van luon
	//	con neu khong phai admin thi minh phai kiem tra xem usr co trong room hay khong
	// laay uuid trong token
	uuid := c.MustGet("uuid").(string)
	roomId := c.Param("roomId")
	result, err := r.roomSer.GetUserInRoomSer(c, roomId, uuid, role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": errors.New("data error"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}
func (r *roomHandler) UpdateRoomHand(c *gin.Context) {
	var update updateRoom
	err := c.ShouldBind(&update)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": "error json ",
		})
		return
	}

	uuidUser := c.MustGet("uuid").(string)
	log.Printf("fomat : %T \t  , %v ", uuidUser, uuidUser)
	role := c.MustGet("role").(string)
	log.Printf("role_system: %v ", c.MustGet("role").(string))
	uuidRoom := c.Param("roomId")
	err = r.roomSer.UpdateRoomSer(c, uuidRoom, update, role, uuidUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": err,
		})
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": "success",
	})

}
