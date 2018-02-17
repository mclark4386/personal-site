package actions

import (
	"fmt"

	"github.com/gobuffalo/buffalo"
	"github.com/gorilla/websocket"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	if websocket.IsWebSocketUpgrade(c.Request()) { //Are they asking for a websocket connection
		conn, err := c.Websocket()
		if err != nil {
			fmt.Println("ERROR creating websocket connection:", err)
		} else {
			go echo(conn)
		}
		return nil
	}

	return c.Render(200, r.HTML("index.html"))
}

func EchoHandler(c buffalo.Context) error {
	conn, err := c.Websocket()
	if err != nil {
		fmt.Println("ERROR creating websocket connection:", err)
	} else {
		go echo(conn)
	}
	return nil
}

func echo(c *websocket.Conn) {
	defer c.Close()

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			fmt.Println("read ws err:", err)
			break
		}

		fmt.Printf("recv: %s\n", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			fmt.Println("write ws err:", err)
			break
		}
	}
}
