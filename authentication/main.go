package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/RipulHandoo/blogx/authentication/pkg/handler/server"
	"github.com/RipulHandoo/blogx/authentication/pkg/routes"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	// get the port 
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error Loading .env file " + err.Error())
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT not set in .env file")
	}

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

	// setup the routes
	v1Router := chi.NewRouter()

	v1Router.Get("/health", server.HealthCheck) // Define the health endpoint here
	router.Mount("/v1", v1Router) // Mount the v1Router on the main router


	apiRouter := routes.SetAllRouter()
	v1Router.Mount("/api",apiRouter)


	fmt.Println("Authentication Server is running on port " + port)
	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}