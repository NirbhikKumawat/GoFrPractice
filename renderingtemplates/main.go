package main

import (
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/http/response"
)

type Todo struct {
	Title string `json:"title"`
	Done  bool   `json:"done"`
}
type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func listHandler(ctx *gofr.Context) (any, error) {
	data := TodoPageData{
		PageTitle: "My TODO list",
		Todos: []Todo{
			{Title: "Expand on Gofr documentation ", Done: false},
			{Title: "Add more  examples", Done: true},
			{Title: "Write some articles", Done: false},
		},
	}
	return response.Template{Data: data, Name: "index.html"}, nil
}

func main() {
	app := gofr.New()
	app.GET("/list", listHandler)
	app.Run()
}
