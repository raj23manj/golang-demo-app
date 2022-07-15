package dao

import (
	"fmt"

	"github.com/raj23manj/demo-app-golang/domain"
	"github.com/raj23manj/demo-app-golang/dto"
)

type userDao struct{}

type userDaoInterface interface {
	GetUser(userId int64) (*domain.User, *dto.DtoErrorResponse)
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

func (u *userDao) GetUser(userId int64) (*domain.User, *dto.DtoErrorResponse) {
	user := users[userId]
	if user == nil {
		return nil, &dto.DtoErrorResponse{
			Message: fmt.Sprintf("user %v does not exists", userId),
		}
	}
	return user, nil
}
