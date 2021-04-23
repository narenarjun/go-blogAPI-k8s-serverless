package controllers

import (
	"blogapi/api/auth"
	"blogapi/api/models"
	"blogapi/api/responses"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// user := models.User{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("we cant read")
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		fmt.Println("we cant unmarshal")
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		fmt.Println("we cant validate")
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := auth.SignIn(user.Email, user.Password)
	if err != nil {
		fmt.Println("we cant signin")
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	responses.JSON(w, http.StatusOK, token)
}
