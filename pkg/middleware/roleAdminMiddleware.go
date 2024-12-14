package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"video-call-project/internal/room"
	"video-call-project/internal/user"
)

// chuyen vao  truy van den csdl cua room de lay quyen cua user  in room
func RoleAdminMiddleware(roomRepo room.IRoomRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		//  lay role-system in token
		roleSystem := c.MustGet("role")
		uuidUser := c.MustGet("uuid").(string)
		roomId := c.Param("roomId")
		if roleSystem == user.Admin {
			c.Next()
		}
		//  lay quyen user in room
		roleUserInRoom, exists, _ := roomRepo.CheckExistsRoomRepo(c, roomId, uuidUser)
		if !exists {
			c.JSON(http.StatusBadRequest, gin.H{
				"role": "user do not exists in room ",
			})
			c.Abort()
			return
		}
		log.Printf("user in room Role : %v", roleUserInRoom.Role)
		if roleUserInRoom.Role == user.Admin {
			c.Next()
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"role": "user do not exists in room ",
			})
			c.Abort()
			return
		}
	}
}
