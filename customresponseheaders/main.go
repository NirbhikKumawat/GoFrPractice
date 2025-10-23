package main

import (
	"time"

	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/http/response"
)

func main() {
	app := gofr.New()
	app.GET("/hello", Hellohandler)

	app.Run()
}
func Hellohandler(ctx *gofr.Context) (any, error) {
	name := ctx.Param("name")
	if name == "" {
		ctx.Log("Name parameter is empty,defaulting to 'World'")
		name = "World"
	}
	headers := map[string]string{
		"X-Custom-Header":  "CustomValue",
		"X-Another-Header": "AnotherValue",
	}

	metaData := map[string]any{
		"environment": "staging",
		"timestamp":   time.Now(),
	}

	return response.Response{
		Data:     map[string]string{"message": "Hello, " + name + "!"},
		Metadata: metaData,
		Headers:  headers,
	}, nil
}
