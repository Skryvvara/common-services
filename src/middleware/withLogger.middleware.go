package middleware

import (
	"log"
	"net/http"

	"github.com/Skryvvara/common-services/config"
	"github.com/go-chi/chi/v5/middleware"
)

var WithLogger func(next http.Handler) http.Handler = middleware.RequestLogger(&middleware.DefaultLogFormatter{
	Logger:  log.Default(),
	NoColor: !config.App.Log.Color,
})
