package api

import (
	"net/http"

	"github.com/mgenov/myproject/domain"
)

func CheckAuth(store domain.SessionStore) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// use store to check session from cookie
			h.ServeHTTP(w, r)
		})
	}
}
