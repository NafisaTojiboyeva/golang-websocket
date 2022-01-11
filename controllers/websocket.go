package controllers

import (
	"log"
	"net/http"
	"strings"

	ws "github.com/gorilla/websocket"
)

type Message struct {
	Data string `json:"data"`
}

var upgrader = ws.Upgrader {
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func (r *http.Request) bool {
		log.Println("New Connection")
		return true
	},
}

func WebSocket(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {

	}

	defer conn.Close()

	for {

		message := Message {}

		err := conn.ReadJSON(&message)

		if err != nil {
			log.Fatal(err)
			break
		}

		err = conn.WriteMessage(1, []byte(strings.ToUpper(message.Data)))

		if err != nil {
			log.Fatal(err)
			break
		}
	}
}