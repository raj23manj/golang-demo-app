package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/raj23manj/demo-app-golang/services"
	"github.com/raj23manj/demo-app-golang/utils/controller"
	"github.com/raj23manj/demo-app-golang/utils/errors"
)

func GetUser(c *gin.Context) {
	userIdParam := c.Param("user_id")
	userId, err := (strconv.ParseInt(userIdParam, 10, 64))
	if err != nil {
		apiErr := errors.NewBadRequestError("user id must be a number!!!")
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
