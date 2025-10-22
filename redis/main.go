package main

import (
	"errors"

	"github.com/redis/go-redis/v9"
	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()
	app.GET("/redis", func(ctx *gofr.Context) (any, error) {
		val, err := ctx.Redis.Get(ctx.Context, "greeting").Result()
		if err != nil && errors.Is(err, redis.Nil) {
			return nil, err
		}
		return val, nil
	})

	app.Run()
}
