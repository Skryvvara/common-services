package routes

import (
	"github.com/Skryvvara/common-services/controllers"
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r *chi.Mux) {
	r.Route("/api", func(r chi.Router) {
		r.Post("/mail", controllers.SendMail)
		r.Get("/coffee", controllers.BrewCoffee)
	})
}
