package router

import (
	"github.com/go-chi/chi"
	"gateway/internal/api/rest/handler"
)

func Route(chiRouter *chi.Mux)  {

	chiRouter.Post("/register", handler.HandleUserRegister)
	chiRouter.Get("/jwt", handler.GenerateJwtToken)

}
