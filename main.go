package main

import (
	"github.com/gohttp/app"
	"github.com/gohttp/response"
	"net/http"
)

func main() {
	app := app.New()
	app.Get("/", func(w http.ResponseWriter, r *http.Request) {
		response.OK(w, "Hello, World")
	})
	app.Listen(":3000")
}
