package router

import (
	"github.com/go-chi/chi"
	"gateway/internal/api/rest/handler"
)

func Route(chiRouter chi.Router)  {

	chiRouter.Post("/register", handler.HandleUserRegister)
	chiRouter.Get("/jwt", handler.GenerateJwtToken)

}
