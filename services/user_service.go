package services

import (
	"net/http"

	"github.com/raj23manj/demo-app-golang/dao"
	"github.com/raj23manj/demo-app-golang/domain"
	"github.com/raj23manj/demo-app-golang/utils/errors"
)

type userService struct{}

type userServiceInterface interface {
	GetUser(userId int64) (*domain.User, errors.ApiError)
}

var (
	UserService userServiceInterface
)

func init() {
	UserService = &userService{}
}

func (u *userService) GetUser(userId int64) (*domain.User, errors.ApiError) {
	response, err := dao.UserDao.GetUser(userId)
	if err != nil {
		return nil, errors.NewApiError(http.StatusNotFound, err.GetMessage())
	}
	return response, nil
}
