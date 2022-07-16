package app

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/raj23manj/demo-app-golang/db/database"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

func StartApp() {
	// testing DB connection
	_, err := database.Database()

	if err != nil {
		log.Println(err)
	}

	mapUrls()

	if err := router.Run("localhost:8080"); err != nil {
		panic(err)
	}
}
