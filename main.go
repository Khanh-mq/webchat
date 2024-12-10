package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"video-call-project/pkg/database"
	"video-call-project/router"
)

func main() {
	var dataMongo = database.ConnectDatabase("user")

	r := gin.Default()

	router.User(r, dataMongo)

	err := r.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

}
