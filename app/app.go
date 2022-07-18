package app

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/raj23manj/demo-app-golang/db/connection"
	"github.com/raj23manj/demo-app-golang/middlewares"
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
	// loggining to a file
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	// middlewares
	router.Use(middlewares.AuthMiddleware())
	// routing urls
	mapUrls()
	// running the server
	if err := router.Run("localhost:8080"); err != nil {
		panic(err)
	}
}
