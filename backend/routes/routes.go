package routes

import (
	"github.com/camilojm27/through-vue-and-go/handlers"

	"github.com/go-chi/chi/v5"
)

func MailRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/all/{page}", handlers.GetMails)
	router.Get("/detail/{id}", handlers.GetMail)
	router.Post("/{param}", handlers.SearchMail)
	return router
}
