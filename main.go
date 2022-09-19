package main

import (
	"github.com/yuan0408/gee"
	"net/http"
)

func main() {
	engine := gee.Default()
	engine.GET("/", func(ctx *gee.Context) {
		ctx.String(http.StatusOK, "Hello yuan\n")
	})

	engine.GET("/panic", func(ctx *gee.Context) {
		names := []string{"yuan"}
		ctx.String(http.StatusOK, names[100])
	})

	engine.Run(":9999")
}
