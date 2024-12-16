package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"video-call-project/pkg/database"
	"video-call-project/router"
)

func main() {
	var dataMongoUser = database.ConnectDatabase("user")
	var dataMongoRoom = database.ConnectDatabase("rooms")
	var dataMongoChat = database.ConnectDatabase("message")

	r := gin.Default()

	router.User(r, dataMongoUser)
	router.Room(r, dataMongoRoom)
	router.ChatRouter(r, dataMongoChat, dataMongoRoom)

	err := r.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

}
