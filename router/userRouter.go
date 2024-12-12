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
	r.POST("/register", newHand.RegisterHandler)

	////  xac thuc
	//// input :email , password
	////output token
	////  sau khi lay  dang nhap xong thi chuyen huong nguoi dung den profile cua ho
	r.POST("/login", newHand.LoginHandler)
	//  sau khi dang nhap xong thi chuyen huong qua user
	//
	//
	//// dang xuat
	//r.POST("/user/logout")
	//
	//// quanly nguoi dung
	//
	u := r.Group("/user", middleware.AuthMiddleware(), middleware.RoleMiddleware(user.Admin, user.Member))
	{
		////  infomation
		////output :id , name , email
		u.GET("/info", newHand.GetUseByIdHandler)
		//
		//// cap nhat thong tin ca nhan
		////input : name , email , .....
		u.PUT("/profile", newHand.UpdateByIdHandler)
		//
		////password thay doi mat khau
		////input : password old , password new
		u.PUT("/change-password", newHand.ChangePasswordHandler)
	}
	//  su dung cho admin
	ad := r.Group("/ad", middleware.AuthMiddleware(), middleware.RoleMiddleware(user.Admin))
	{
		//wuan ly tai khoan nguoi dung
		ad.GET("/user", newHand.GetUserAll)
		// lay thong tin nguoi dung cu the
		ad.GET("/user/:uuid", newHand.GetUserById)
		//  co the tao tai khoan moi cho user
		ad.POST("/register", newHand.RegisterHandler)
		ad.PUT("/user/:uuid", newHand.UpdateByIdHandler)

		// quan ly quan  quyen  role

		ad.PUT("/user/:uuid/role", newHand.RoleRightsHandler)
		//  lay danh sach cac quyen trong he thong theo role
		ad.GET("user/role")
	}

}
