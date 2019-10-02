package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

func main() {

	header := make(http.Header)
	header.Add("Origin", "http://localhost:1234")
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:1234/ws", header)

	if err != nil {
		fmt.Println("toto", err.Error())
		os.Exit(1)
	}

	for {
		_, reply, err := conn.ReadMessage()

		if err != nil {
			if err == io.EOF {
				fmt.Println(`EOF from server`)
				break
			}

			if websocket.IsCloseError(err, websocket.CloseAbnormalClosure) {
				fmt.Println("Close from server")
				break
			}
		}
		fmt.Println("Received from server: ", string(reply[:]))

		err = conn.WriteMessage(websocket.TextMessage, reply)
		if err != nil {
			fmt.Println("Could not return msg")
			break
		}

	}

	os.Exit(0)
}
