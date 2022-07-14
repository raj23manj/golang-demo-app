package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/raj23manj/demo-app-golang/domain/utils"
	"github.com/raj23manj/demo-app-golang/services"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	userIdParam := req.URL.Query().Get("user_id")
	userId, err := (strconv.ParseInt(userIdParam, 10, 64))
	if err != nil {
		apiError := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonValue, _ := json.Marshal(apiError)
		resp.WriteHeader(apiError.StatusCode)
		resp.Write([]byte(jsonValue))
		return
	}

	user, apiError := services.GetUser(userId)
	if apiError != nil {
		jsonValue, _ := json.Marshal(apiError)
		resp.WriteHeader(apiError.StatusCode)
		resp.Write([]byte(jsonValue))
		return
	}

	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}
