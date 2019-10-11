package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
)

// https://github.com/gin-gonic/gin/issues/124
var Socketio_Server *socketio.Server

func socketHandler(c *gin.Context) {
	Socketio_Server.On("connection", func(so socketio.Socket) {
		fmt.Println("on connection")

		so.Join("chat")

		so.On("chat message", func(msg string) {
			fmt.Println("emit:", so.Emit("chat message", msg))
			so.BroadcastTo("chat", "chat message", msg)
		})
		so.On("disconnection", func() {
			fmt.Println("on disconnect")
		})
	})

	Socketio_Server.On("error", func(so socketio.Socket, err error) {
		fmt.Printf("[ WebSocket ] Error : %v", err.Error())
	})

	Socketio_Server.ServeHTTP(c.Writer, c.Request)
}

func main() {
	var router_engine = gin.Default()
	var err error

	Socketio_Server, err = socketio.NewServer(nil)
	if err != nil {
		panic(err)
	}

	router_engine.GET("/", IndexHandler)
	router_engine.Static("/public", "./public")

	router_engine.GET("/socket.io", socketHandler)
	router_engine.POST("/socket.io", socketHandler)
	router_engine.Handle("WS", "/socket.io", []gin.HandlerFunc{socketHandler})
	router_engine.Handle("WSS", "/socket.io", []gin.HandlerFunc{socketHandler})

	router_engine.Run(":8000")
}
