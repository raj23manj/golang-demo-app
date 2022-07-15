package dao

import (
	"fmt"
	"net/http"

	"github.com/raj23manj/demo-app-golang/domain"
	"github.com/raj23manj/demo-app-golang/utils/errors"
)

type userDao struct{}

type userDaoInterface interface {
	GetUser(userId int64) (*domain.User, *errors.ApplicationError)
}

var (
	users = map[int64]*domain.User{
		123: &domain.User{Id: 123, FirstName: "Fede", LastName: "Leon", Email: "myemail@gmail.com"},
	}

	UserDao userDaoInterface
)

func init() {
	UserDao = &userDao{}
}

func (u *userDao) GetUser(userId int64) (*domain.User, *errors.ApplicationError) {
	user := users[userId]
	if user == nil {
		return nil, &errors.ApplicationError{
			Message:    fmt.Sprintf("user %v does not exists", userId),
			StatusCode: http.StatusNotFound,
			Code:       "not_found",
		}
	}
	return user, nil
}
