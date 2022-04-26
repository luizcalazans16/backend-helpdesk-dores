package controller

import (
	"backend-helpdesk-dores/src/response"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := ioutil.ReadAll(r.body)

	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
	}

	var user model.User

	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		response.Error(w, http.StatusBadRequest, err)
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repository.NewRepositoryOfUser(db)
	useOfDatabase, err := repository.SearchEmail(user.Email)
}
