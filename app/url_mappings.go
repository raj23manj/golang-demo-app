package app

import "github.com/raj23manj/demo-app-golang/controllers"

func mapUrls() {
	router.GET("/users/:user_id", controllers.GetUser)
}
