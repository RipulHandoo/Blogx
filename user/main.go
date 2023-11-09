package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/RipulHandoo/blogx/user/pkg/handler/server"
	"github.com/RipulHandoo/blogx/user/pkg/routers"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Printf("Error loading .env file in user/main.go %v", err)
		return
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not defined in the user/.env file")
		return
	}

	// set up router
	router := chi.NewRouter()

	// Configure CORS (Cross-Origin Resource Sharing) settings.
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Links"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/health", server.HealthCheck)
	router.Mount("/v1", v1Router)

	// set the router for api calls
	apiRouter := routers.SetAllRouter()
	v1Router.Mount("/api", apiRouter)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	fmt.Print("User Server running on http://localhost: " + port)

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting server: %v", err)
	}
}
