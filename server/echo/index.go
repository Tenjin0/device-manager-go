package main

import (
	"fmt"
	"log"

	socketio "github.com/googollee/go-socket.io"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	e.Static("/assets", "../../assets")
	e.File("/", "../../assets/index.html")
	e.Any("/socket.io/", SocketIOWrapper(server))

	e.GET("to", func(c echo.Context) error {
		fmt.Println(server.Rooms("/"))
		return c.JSON(200, nil)
	})
	e.Logger.Fatal(e.Start(":1323"))
}

// func ServerHeader(c echo.Context) error {

// 	server, err := socketio.NewServer(nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	server.ServeHTTP(c.Response(), c.Request())
// 	return nil

// }

func SocketIOWrapper(server *socketio.Server) echo.HandlerFunc {

	wrapper, err := NewWrapperWithServer(server)
	if err != nil {
		log.Fatal(err)
	}
	return wrapper.HandlerFunc
}
