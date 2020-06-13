package main

import (
	"fmt"
	"log"

	"github.com/Tenjin0/device-manager-echo/server/socket"
	socketio "github.com/googollee/go-socket.io"
	"github.com/labstack/echo"
)

type CustomContext struct {
	echo.Context
	socket *socketio.Server
}

func main() {
	e := echo.New()

	server, err := socketio.NewServer(nil)

	if err != nil {
		log.Fatal(err)
	}

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &CustomContext{c, server}
			return next(cc)
		}
	})

	e.Static("/assets", "../../assets")
	e.File("/", "../../assets/index.html")
	e.Any("/socket.io/", SocketIOWrapper(server))

	e.GET("to", func(c echo.Context) error {
		fmt.Println(server.Rooms("/"))
		return c.JSON(200, nil)
	})
	e.Logger.Fatal(e.Start(":1323"))
}

	e.GET("to", func(c echo.Context) error {
		fmt.Println(server.Rooms("/"))
		return c.JSON(200, nil)
	})

	e.Logger.Fatal(e.Start(":1323"))
}

func SocketIOWrapper(server *socketio.Server) echo.HandlerFunc {

	wrapper, err := socket.NewWrapperWithServer(server)
	if err != nil {
		log.Fatal(err)
	}

	return wrapper.HandlerFunc
}
