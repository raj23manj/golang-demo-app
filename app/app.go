package app

import (
	"github.com/gin-gonic/gin"
	"github.com/raj23manj/demo-app-golang/db/connection"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

func StartApp() {
	// DB connection init
	connection.Database()
	// routing urls
	mapUrls()
	// running the server
	if err := router.Run("localhost:8080"); err != nil {
		panic(err)
	}
}
