package middleware

import (
	"net/http"
	"regexp"
)

type CorsConfig struct {
	Enabled         bool
	WhiteListRegexp string
}

func allowedOrigin(origin string, c *CorsConfig) bool {
	ok, _ := regexp.MatchString(c.WhiteListRegexp, origin)
	return ok && c.Enabled
}

func CorsMiddleware(handler http.Handler, c *CorsConfig) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")

		origin := r.Header.Get("Origin")
		if allowedOrigin(origin, c) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
		}

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		handler.ServeHTTP(w, r)
	})
}
