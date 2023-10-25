package routes

import (
	"github.com/RipulHandoo/blogx/authentication/pkg/handler/server"
	"github.com/go-chi/chi"
)

func SetAllRouter() chi.Router {
	apiRouter := chi.NewRouter()

	apiRouter.Get("/",server.HealthCheck)

	authRouter := SetAuthRouter()
	apiRouter.Mount("/auth",authRouter)

	return apiRouter
}