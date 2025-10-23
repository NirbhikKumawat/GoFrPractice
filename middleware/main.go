package main

import (
	"net/http"

	"gofr.dev/pkg/gofr"
	gofrHTTP "gofr.dev/pkg/gofr/http"
)

func customMiddleware() gofrHTTP.Middleware {
	return func(inner http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// custom logic

			inner.ServeHTTP(w, r)
		})
	}
}

func main() {
	app := gofr.New()
	app.UseMiddleware(customMiddleware())
	app.Run()
}
