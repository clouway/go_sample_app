package api

import (
	"fmt"
	"net/http"
)

func IndexHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Hello !")
	})
}
