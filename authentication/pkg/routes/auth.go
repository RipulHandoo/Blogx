package routes

import (
	auth "github.com/RipulHandoo/blogx/authentication/pkg/handler/auth"
	"github.com/RipulHandoo/blogx/authentication/pkg/handler/server"
	"github.com/RipulHandoo/blogx/authentication/pkg/middleware"
	"github.com/go-chi/chi"
)

func SetAuthRouter() chi.Router {
	authRouter := chi.NewRouter()
	authRouter.Get("/",server.HealthCheck)
	authRouter.Post("/singUp", auth.HandleRegisterUser)
	authRouter.Post("/login",auth.HandlerLoginuser)
	authRouter.Post("/logout", middleware.Auth(middleware.AuthHandler(auth.HandleUserLogout)))

	return authRouter
}