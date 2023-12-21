package main

import (
	"log"
	"net/http"

	"nhooyr.io/websocket"
)

func main() {
	http.HandleFunc("/ws/echo", func(w http.ResponseWriter, r *http.Request) {
		conn, err := websocket.Accept(w, r, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		defer conn.Close(websocket.StatusNormalClosure, "")

		ctx := r.Context()
		for {
			mt, b, err := conn.Read(ctx)
			if err != nil {
				log.Println(err)
				return
			}

			log.Println(string(b))

			err = conn.Write(ctx, mt, b)
			if err != nil {
				log.Println(err)
				return
			}
		}
	})
	http.ListenAndServe(":9090", nil)
}
