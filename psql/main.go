package main

import (
	"errors"

	"github.com/redis/go-redis/v9"
	"gofr.dev/pkg/gofr"
)

type Customer struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	app := gofr.New()

	app.GET("/redis", func(ctx *gofr.Context) (any, error) {
		val, err := ctx.Redis.Get(ctx.Context, "test").Result()
		if err != nil && errors.Is(err, redis.Nil) {
			return nil, err
		}
		return val, nil
	})
	app.POST("/customer/{name}", func(ctx *gofr.Context) (any, error) {
		name := ctx.PathParam("name")
		_, err := ctx.SQL.ExecContext(ctx.Context, "INSERT INTO customers (name) VALUES ($1)", name)
		return nil, err
	})
	app.GET("/customer", func(ctx *gofr.Context) (any, error) {
		var customers []Customer
		rows, err := ctx.SQL.QueryContext(ctx, "SELECT * FROM customers")
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			var customer Customer
			if err := rows.Scan(&customer.ID, &customer.Name); err != nil {
				return nil, err
			}
			customers = append(customers, customer)
		}
		return customers, nil
	})
	app.Run()
}
