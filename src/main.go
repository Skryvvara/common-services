package main

import (
	"log"
	"net/http"

	"github.com/Skryvvara/common-services/config"
	"github.com/Skryvvara/common-services/logger"
	"github.com/Skryvvara/common-services/middleware"
	"github.com/Skryvvara/common-services/routes"
	"github.com/go-chi/chi/v5"
)

func main() {
	config.Initialize()
	logger.Initialize()

	r := chi.NewRouter()
	r.Use(middleware.WithLogger)
	r.Use(middleware.WithHeaders)

	routes.RegisterRoutes(r)

	port := config.App.Server.Port
	log.Println("--- COMMON SERVICES API ---")
	log.Println("--- listening on port: " + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
