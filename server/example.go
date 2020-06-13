// package main

// import (
// 	"fmt"
// 	"log"

// 	socketio "github.com/googollee/go-socket.io"
// 	"github.com/labstack/echo"
// 	esi "github.com/umirode/echo-socket.io"
// )

// func main() {
// 	e := echo.New()

// 	server, err := socketio.NewServer(nil)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	e.Static("/assets", "../../assets")
// 	e.File("/", "../../assets/index.html")
// 	e.Any("/socket.io/", socketIOWrapper(server))
// 	e.GET("/to")
// 	e.Logger.Fatal(e.Start(":1323"))
// }

// func socketIOWrapper(server *socketio.Server) func(context echo.Context) error {
// 	wrapper, err := esi.NewWrapperWithServer(server)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return nil
// 	}

// 	// must be implemented to work
// 	wrapper.OnConnect("", func(context echo.Context, conn socketio.Conn) error {
// 		// Must be set to work
// 		conn.SetContext("")
// 		fmt.Println("connected:", conn.ID())
// 		return nil
// 	})

// 	return func(context echo.Context) error {
// 		go wrapper.Server.Serve()

// 		wrapper.Context = context
// 		wrapper.Server.ServeHTTP(context.Response(), context.Request())
// 		return nil
// 	}
// }
