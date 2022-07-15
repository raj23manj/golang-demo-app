package services

import (
	"github.com/raj23manj/demo-app-golang/dao"
	"github.com/raj23manj/demo-app-golang/domain"
	"github.com/raj23manj/demo-app-golang/utils/errors"
)

type userService struct{}

type userServiceInterface interface {
	GetUser(userId int64) (*domain.User, *errors.ApplicationError)
}

var (
	UserService userServiceInterface
)

func init() {
	UserService = &userService{}
}

func (u *userService) GetUser(userId int64) (*domain.User, *errors.ApplicationError) {
	return dao.UserDao.GetUser(userId)
}
