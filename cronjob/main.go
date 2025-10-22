package main

import (
	"time"

	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()
	app.AddCronJob("*/10 * * * * *", "", func(ctx *gofr.Context) {
		ctx.Logger.Infof("Current time is %v", time.Now())
	})
	app.Run()
}
