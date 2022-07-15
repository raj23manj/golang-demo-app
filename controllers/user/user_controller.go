package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/raj23manj/demo-app-golang/services"
	"github.com/raj23manj/demo-app-golang/utils/app_errors"
	"github.com/raj23manj/demo-app-golang/utils/controller"
)

func GetUser(c *gin.Context) {
	userIdParam := c.Param("user_id")
	userId, err := (strconv.ParseInt(userIdParam, 10, 64))
	if err != nil {
		apiErr := &app_errors.ApplicationError{
			Message:    "user_id must be a number$$$",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		controller.RespondError(c, apiErr)
		return
	}

	user, apiErr := services.UserService.GetUser(userId)
	if apiErr != nil {
		controller.RespondError(c, apiErr)
		return
	}

	controller.Respond(c, http.StatusOK, user)
}
