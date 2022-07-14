package services

import (
	"github.com/raj23manj/demo-app-golang/domain"
	"github.com/raj23manj/demo-app-golang/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
