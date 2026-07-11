package router

import (
	"log"
	"gateway/internal/api/rest/handler"
	"net/http"
	"github.com/go-chi/chi"
)

func Route() *chi.Mux {
	log.Println("req is in router layer")
	chiRouter := chi.NewRouter()

	chiRouter.Get("/register", func(w http.ResponseWriter, r *http.Request) {
		handler.HandleRegister(w, r)
	})

	return chiRouter
}
