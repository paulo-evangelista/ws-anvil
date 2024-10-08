package main

import (
	"fmt"
	"github.com/gorilla/websocket"
)

func messageLoop() {
	fmt.Println("entered message loop")

	for {
		msg := <-incomingLogs
		for client := range connectedClients {
			client.WriteMessage(websocket.TextMessage, []byte(msg))
		}
	}
}
