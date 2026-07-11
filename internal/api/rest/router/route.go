package router

import (
	"gateway/internal/api/rest/handler"
	"net/http"
	"github.com/go-chi/chi"
)

func Route(chiRouter *chi.Mux)  {

	chiRouter.Get("/register", func(w http.ResponseWriter, r *http.Request) {
		handler.HandleRegister(w, r)
	})

}
