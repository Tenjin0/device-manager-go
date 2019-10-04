package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	gosocketio "github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
)

// var wsupgrader = websocket.Upgrader{
// 	ReadBufferSize:  1024,
// 	WriteBufferSize: 1024,
// 	CheckOrigin: func(r *http.Request) bool {
// 		return true
// 	},
// }

func wshandler() *gosocketio.Server {
	server := gosocketio.NewServer(transport.GetDefaultWebsocketTransport())

	//handle connected
	server.On(gosocketio.OnConnection, func(c *gosocketio.Channel) {
		log.Println("New client connected")
		//join them to room
		c.Join("chat")
	})

	type Message struct {
		Name    string `json:"name"`
		Message string `json:"message"`
	}

	//handle custom event
	server.On("send", func(c *gosocketio.Channel, msg Message) string {
		fmt.Println(msg)
		//send event to all in room
		return "OK"
	})

	//setup http server

	return server
}

func main() {

	r := gin.Default()
	// server := wshandler()

	r.StaticFile("/", "./asset")

	r.GET("/socket.io/", func(c *gin.Context) {
		c.String(200, "Hello")
	})

	r.Run("localhost:1234")
}
