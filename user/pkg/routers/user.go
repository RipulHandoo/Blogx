package routers

import (
	"github.com/RipulHandoo/blogx/user/pkg/handler/server"
	users "github.com/RipulHandoo/blogx/user/pkg/handler/user"
	"github.com/RipulHandoo/blogx/user/pkg/middleware"
	"github.com/go-chi/chi"
)

func SetUserRoutes() chi.Router {
	userRouter := chi.NewRouter()

	userRouter.Get("/", server.HealthCheck)
	userRouter.Post("/follow", middleware.Auth(middleware.AuthHandler(users.FollowUser)))
	
	return userRouter
}
