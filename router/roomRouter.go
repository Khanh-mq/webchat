package router

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"video-call-project/pkg/middleware"
)

func Room(r *gin.Engine, db *mongo.Collection) {
	// user
	user := r.Group("/room", middleware.AuthMiddleware())
	{
		//  lay danh sach cac
		user.GET("")

	}
	// admin

}
