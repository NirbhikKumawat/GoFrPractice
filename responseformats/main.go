package main

import (
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/http/response"
)

type user struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	app := gofr.New()
	app.GET("/users1", func(ctx *gofr.Context) (any, error) {
		users := []user{{ID: 1, Name: "alice"}, {ID: 2, Name: "bob"}}
		return users, nil
	})
	app.GET("/users2", func(ctx *gofr.Context) (any, error) {
		users := []user{{ID: 1, Name: "alice"}, {ID: 2, Name: "bob"}}
		return response.Raw{Data: users}, nil
	})
	app.Run()

}
