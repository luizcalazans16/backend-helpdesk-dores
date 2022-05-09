package controller

import (
	"backend-helpdesk-dores/internal/config/authentication"
	"backend-helpdesk-dores/internal/config/security"
	"backend-helpdesk-dores/internal/core/model"
	"backend-helpdesk-dores/internal/core/model/response"
	"backend-helpdesk-dores/internal/core/service"
	"backend-helpdesk-dores/internal/database"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var loginService *service.LoginService
	bodyRequest, err := ioutil.ReadAll(r.Body)

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

	useOfDatabase, err := loginService.SearchEmail(user.Email)

	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.VerifyPassword(useOfDatabase.Password, user.Password); err != nil {
		response.Error(w, http.StatusUnauthorized, err)
	}

	token, error := authentication.GenerateToken(2)
	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}
	fmt.Println(token)
	response.JSON(w, http.StatusOK, nil)
}
