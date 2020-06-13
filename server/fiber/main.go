package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber"
	"github.com/gofiber/logger"
	socketio "github.com/googollee/go-socket.io"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

func WrapHandler(f func(http.ResponseWriter, *http.Request)) func(ctx *fiber.Ctx) {
	return func(ctx *fiber.Ctx) {
		fasthttpadaptor.NewFastHTTPHandler(http.HandlerFunc(f))(ctx.Fasthttp)
	}
}

func main() {

	app := fiber.New()
	app.Use(logger.New())
	app.Static("/", "../../assets/index.html")
	app.Static("/assets", "./assets")

	server, err := socketio.NewServer(nil)

	if err != nil {
		log.Fatal(err)
	}

	app.All("/socket.io/*", WrapHandler(server.ServeHTTP))

	app.Listen(1323)
}
