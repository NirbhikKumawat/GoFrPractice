package main

import "gofr.dev/pkg/gofr"

func main() {
	app := gofr.New()
	app.EnableAPIKeyAuth("9221e451-451f-4cd6-a23d-2b2d3adea9cf", "0d98ecfe-4677-48aa-b463-d43505766915")

	app.GET("/customer", func(ctx *gofr.Context) (any, error) {
		return "api-key authenticated", nil
	})

	app.Run()
}
