package routers

import (
	"github.com/RipulHandoo/blogx/user/pkg/handler/server"
	"github.com/go-chi/chi"
)

func SetAllRouter() chi.Router {
	apiRouter := chi.NewRouter()

	// set the router for api calls
	apiRouter.Get("/", server.HealthCheck)

	userRouter := SetUserRoutes()
	apiRouter.Mount("/user", userRouter)

	return apiRouter
}
