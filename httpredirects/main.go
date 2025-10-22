package main

import (
	"gofr.dev/pkg/gofr"

	"gofr.dev/pkg/gofr/http/response"
)

func main() {
	app := gofr.New()
	app.GET("/old-page", func(ctx *gofr.Context) (any, error) {
		return response.Redirect{URL: "https://example.com/new-page"}, nil
	})
	app.Run()
}
