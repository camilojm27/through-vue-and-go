package handlers

import (
	"encoding/json"
	"github.com/camilojm27/through-vue-and-go/utils"
	"net/http"
)

type User struct {
	Name string `json:"name"`
}

var users []User

func GetMails(w http.ResponseWriter, r *http.Request) {
	//MakeAPIRequest("")
	request, err := utils.MakeAPIRequest()
	if err != nil {
		return
	}

	w.Write(request)
}

func SearchMail(w http.ResponseWriter, r *http.Request) {
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	users = append(users, user)
	json.NewEncoder(w).Encode(user)
}
