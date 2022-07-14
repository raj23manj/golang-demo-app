package app

import "github.com/raj23manj/demo-app-golang/controllers/user"

func mapUrls() {
	router.GET("/users/:user_id", user.GetUser)
}
