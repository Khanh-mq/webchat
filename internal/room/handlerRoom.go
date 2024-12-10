package room

import (
	"fmt"
	"github.com/gin-gonic/gin"
	gguid "github.com/google/uuid"
	"github.com/gorilla/websocket"

	"net/http"
)

var (
	upgrade = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func RoomCreate(c *gin.Context) {
	uuid := gguid.New().String()
	c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/room/%s", uuid))
}
func Room(c *gin.Context) {
	uuid := c.Param("uuid")
	if uuid == "" {
		c.Status(400)
		return
	}
	uuid, suuid, _ := createOrGetRoom(uuid)
}

func createOrGetRoom(uuid string) (string, string, Room) {

}
func RoomWebSocket(c *gin.Context) {
	uuid := c.Param("uuid")
	if uuid == "" {
		c.Status(400)
		return
	}
	_, _, room := createOrGetRoom(uuid)
}
