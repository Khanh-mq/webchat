package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"video-call-project/internal/room"
)

func CheckUserInRoomMiddleware(room room.IRoomRepo) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		roomId := ctx.Param("roomId")
		uuid := ctx.MustGet("uuid").(string)
		log.Printf("CheckUserInRoomMiddleware roomId:%s uuid:%s", roomId, uuid)
		// Lấy danh sách user trong room
		users, err := room.GetUserInRoomRepo(ctx, roomId)
		if err != nil {
			log.Println("Error fetching users in room:", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code": "error",
				"msg":  "Failed to fetch users in room",
			})
			ctx.Abort()
			return
		}

		// Kiểm tra nếu user có tồn tại trong room
		for _, user := range users {
			if user.UserId == uuid {
				// User tồn tại, cho phép tiếp tục
				ctx.Set("userRole", user) // Lưu thông tin user để sử dụng sau
				ctx.Next()
				return
			}
		}

		// User không tồn tại trong room
		ctx.JSON(http.StatusForbidden, gin.H{
			"code": "user_not_in_room",
			"msg":  "User does not exist in this room",
		})
		ctx.Abort()
	}
}
