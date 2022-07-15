package services

import (
	"github.com/raj23manj/demo-app-golang/domain"
	"github.com/raj23manj/demo-app-golang/utils"
)

type userService struct{}

type userServiceInterface interface {
	GetUser(userId int64) (*domain.User, *utils.ApplicationError)
}

var (
	UserService userServiceInterface
)

func init() {
	UserService = &userService{}
}

func (u *userService) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
