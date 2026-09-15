package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader websocket.Upgrader

var connections = make([]*websocket.Conn, 0)

func doc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello")

	conn, err := upgrader.Upgrade(w, r, nil)
	connections = append(connections, conn)
	if err != nil {
		fmt.Println("upgrade error:", err)
		return
	}
	defer conn.Close()

	for {
		_, m, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("read error:", err)
			return
		}

		for _, r := range connections {
			r.WriteJSON(map[string]string{
				"typed": string(m),
			})
		}

		fmt.Println(string(m))
	}
}

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello server is running")
	})
	http.HandleFunc("/doc", doc)
	http.ListenAndServe(":8080", nil)
}
