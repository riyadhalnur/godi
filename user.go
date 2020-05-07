package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/simple-go-server/utils"
)

type User struct {
	ID    string   `json:"id,omitempty"`
	Name  string   `json:"name,omitempty"`
	Roles []string `json:"roles,omitempty"`
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	pathVars := mux.Vars(r)
	log.Println(pathVars)
	if pathVars["id"] != "" {
		user, _ := json.Marshal(User{
			ID:    pathVars["id"],
			Name:  "Test User",
			Roles: []string{"admin"},
		})

		utils.RespondJSON(w, &utils.Response{
			StatusCode: http.StatusOK,
			Body:       user,
		})
	}

	return
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errMsg, _ := json.Marshal(utils.ErrorResponse{
			Code:    1,
			Type:    "Server Error",
			Message: err.Error(),
		})

		utils.RespondJSON(w, &utils.Response{
			StatusCode: http.StatusInternalServerError,
			Body:       errMsg,
		})
		return
	}

	var user User
	if err := json.Unmarshal(body, &user); err != nil {
		errMsg, _ := json.Marshal(utils.ErrorResponse{
			Code:    1,
			Type:    "Invalid arguments",
			Message: err.Error(),
		})

		utils.RespondJSON(w, &utils.Response{
			StatusCode: http.StatusBadRequest,
			Body:       errMsg,
		})
		return
	}

	utils.RespondJSON(w, &utils.Response{
		StatusCode: http.StatusAccepted,
		Body:       body,
	})

	return
}
