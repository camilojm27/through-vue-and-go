package handlers

import (
	"github.com/go-chi/chi/v5"
	"net/http"

	"github.com/camilojm27/through-vue-and-go/utils"
)

func GetMails(w http.ResponseWriter, r *http.Request) {
	//MakeAPIRequest("")
	page := chi.URLParam(r, "page")
	response, err := utils.MakeAPIRequest(page, "", "")
	if err != nil {
		return
	}
	w.Write(response)
}

func SearchMail(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "param")

	response, err := utils.MakeAPIRequest("", param, "")
	if err != nil {
		return
	}
	w.Write(response)
}

func GetMail(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "id")

	response, err := utils.MakeAPIRequest("", "", param)
	if err != nil {
		return
	}
	w.Write(response)
}
