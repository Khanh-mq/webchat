package router

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"video-call-project/internal/user"
	"video-call-project/pkg/middleware"
)

// nếu thích sử dụng reddis thì nên dùng cũng đuọcw
func User(r *gin.Engine, db *mongo.Collection) {

	newRepo := user.NewUserRepository(db)
	newSer := user.NewUserService(newRepo)
	newHand := user.NewUserHandler(*newSer)

	//name , email , password (input)
	// output :message : success
	r.POST("/user/register", newHand.RegisterHandler)

	////  xac thuc
	//// input :email , password
	////output token
	////  sau khi lay  dang nhap xong thi chuyen huong nguoi dung den profile cua ho
	r.POST("/user/login", newHand.LoginHandler)
	//
	//
	//// dang xuat
	//r.POST("/user/logout")
	//
	//// quanly nguoi dung
	//
	////  infomation
	////output :id , name , email
	r.GET("/user/info", middleware.AuthMiddleware(), newHand.GetUseByIdHandler)
	//
	//// cap nhat thong tin ca nhan
	////input : name , email , .....
	r.PUT("/user/profile", middleware.AuthMiddleware(), newHand.UpdateByIdHandler)
	//
	////password thay doi mat khau
	////input : password old , password new
	r.PUT("/user/change-password", middleware.AuthMiddleware(), newHand.ChangePasswordHandler)
}
