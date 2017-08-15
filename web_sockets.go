package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool { return true },
}

func main()  {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":4000", nil)
}

func handler(w http.ResponseWriter, r *http.Request)  {
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		messageType, message, err := socket.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(string(message))

		if err = socket.WriteMessage(messageType, message); err != nil {
			fmt.Println(err)
			return
		}
	}
}

