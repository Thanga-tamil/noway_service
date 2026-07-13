package router

import (
	"gateway/internal/api/rest/handler"
	"net/http"
	"github.com/go-chi/chi"
)

func Route(chiRouter *chi.Mux)  {

	chiRouter.Get("/register", func(w http.ResponseWriter, r *http.Request) {
		handler.HandleUserRegister(w, r)
	})

	chiRouter.Get("/jwt", func(w http.ResponseWriter, r *http.Request) {
		handler.GenerateJwtToken(w, r)
	})
}
