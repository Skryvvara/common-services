package middleware

import (
	"net/http"

	"github.com/Skryvvara/common-services/config"
)

func WithHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("X-Powered-By", config.App.Server.PoweredByHeader)

		next.ServeHTTP(w, r)
	})
}
