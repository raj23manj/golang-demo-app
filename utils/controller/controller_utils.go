package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/raj23manj/demo-app-golang/utils/errors"
)

func Respond(c *gin.Context, status int, body interface{}) {
	if c.GetHeader("Accept") == "application/xml" {
		c.XML(status, body)
		return
	}
	c.JSON(status, body)
}

func RespondError(c *gin.Context, err errors.ApiError) {
	if c.GetHeader("Accept") == "application/xml" {
		c.XML(err.EStatus(), err)
		return
	}
	c.JSON(err.EStatus(), err)
}
