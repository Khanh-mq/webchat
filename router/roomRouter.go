package router

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"video-call-project/internal/room"
	"video-call-project/pkg/middleware"
)

func Room(r *gin.Engine, db *mongo.Collection) {
	// user
	newRepo := room.NewRoomRepo(db)
	newSer := room.NewRoomService(newRepo)
	newHand := room.NewRoomHandler(newSer)
	user := r.Group("/rooms", middleware.AuthMiddleware())
	{
		//  lay danh sach cac
		//  user co the toa phong va nguoi tao phong thi xe la chu phong
		user.POST("/create", newHand.CreateRoomHand)
		//  xem thong tin chi tiet phong
		user.GET("/listRoom", newHand.GetListRoomHand)
		user.POST("/:roomId")
		//tham gia phong
		user.POST("/:roomId/join", newHand.JoinRoomHand)
		// roi khoi phong
		user.POST("/:roomId/leave")
		//xem danh sach thanh vien trong phong
		// viec xem danh sach ban phai co trong room da

		user.GET("/:roomId/users", newHand.GetUserRoomHand)

		// cap nhat thong tin phong  admin-system and aadmin-group

	}
	// admin
	admin := r.Group("/admin", middleware.AuthMiddleware(), middleware.RoleAdminMiddleware(newRepo))
	{
		// cap nhat trong them so nguoi hay doi ten phong  //
		//  day admin hoawxj admin trong room deu co the
		admin.POST("/:roomId/update", newHand.UpdateRoomHand)

		//xoa phong
		admin.POST("/:roomId/delete", newHand.DeletedRoomHand)
		//  them nguoi dung vao phong
		admin.POST("/:roomId/users", newHand.AddUserInRoomHand)

		// xoa nguoi ra khoi phong

		admin.POST("/:roomId/users/:uuid", newHand.DeleteUserInRoomHand)

		//  phan quyen nguoi dung trong room
		admin.POST("/:roomId/users/:uuid/role", newHand.UpdateRoleInRoomHand)

		// khoa mo phong
		//
		admin.POST("/:roomId/lock")

	}

}
