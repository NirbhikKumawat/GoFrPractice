package main

import "gofr.dev/pkg/gofr"

func main() {
	app := gofr.New()
	app.EnableBasicAuth("admin", "secret")
	app.GET("/protected-resource", func(ctx *gofr.Context) (any, error) {
		return "authenticated", nil
	})
	app.Run()
}
