package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/raj23manj/demo-app-golang/services"
	"github.com/raj23manj/demo-app-golang/utils"
)

func GetUser(c *gin.Context) {
	userIdParam := c.Param("user_id")
	userId, err := (strconv.ParseInt(userIdParam, 10, 64))
	if err != nil {
		apiError := &utils.ApplicationError{
			Message:    "user_id must be a number$$$",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		c.JSON(apiError.StatusCode, apiError)
		return
	}

	user, apiError := services.GetUser(userId)
	if apiError != nil {
		c.JSON(apiError.StatusCode, apiError)
		return
	}

	c.JSON(http.StatusOK, user)
}
