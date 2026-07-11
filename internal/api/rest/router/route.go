package router

import (
	"log"
	"gateway/internal/api/rest/handler"
	"net/http"
	"github.com/go-chi/chi"
)

func Route(chiRouter *chi.Mux)  {
	log.Println("req is in router layer")

	chiRouter.Get("/register", func(w http.ResponseWriter, r *http.Request) {
		handler.HandleRegister(w, r)
	})

}
