package main

import (
	"slices"

	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/container"
)

func apiKeyValidator(c *container.Container, apiKey string) bool {
	validKeys := []string{"f0e1dffd-0ff0-4ac8-92a3-22d44a1464e4", "d7e4b46e-5b04-47b2-836c-2c7c91250f40"}
	return slices.Contains(validKeys, apiKey)
}
func main() {
	app := gofr.New()
	app.EnableAPIKeyAuthWithValidator(apiKeyValidator)
	app.GET("/customer", func(ctx *gofr.Context) (any, error) {
		return "Hello World", nil
	})
	app.Run()
}
