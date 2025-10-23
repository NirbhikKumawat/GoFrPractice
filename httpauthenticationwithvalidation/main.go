package main

import (
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/container"
)

func validateUser(c *container.Container, username, password string) bool {
	return username == "john" && password == "doe123"
}

func main() {
	app := gofr.New()
	app.EnableBasicAuthWithValidator(validateUser)
	app.GET("/secure-data", func(c *gofr.Context) (any, error) {
		return nil, nil
	})
	app.Run()
}
