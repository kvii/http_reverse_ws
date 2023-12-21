package main

import (
	"context"
	"fmt"
	"log"

	"nhooyr.io/websocket"
)

func main() {
	ctx := context.Background()

	conn, _, err := websocket.Dial(ctx, "ws://localhost:9090/ws/echo", nil)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close(websocket.StatusNormalClosure, "")

	log.Println("start")

	for {
		var s string
		_, err := fmt.Scan(&s)
		if err != nil {
			return
		}
		switch s {
		case "quit", "q", "exit", "e":
			return
		}

		err = conn.Write(ctx, websocket.MessageText, []byte(s))
		if err != nil {
			log.Println(err)
			return
		}

		mt, b, err := conn.Read(ctx)
		if err != nil {
			log.Println(err)
			return
		}
		switch mt {
		case websocket.MessageText:
			log.Println(string(b))
		}
	}
}
