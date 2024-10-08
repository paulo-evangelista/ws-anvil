package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
	"os/exec"
	"sync"
)

var (
	connectedClients = make(map[*websocket.Conn]bool)
	anvilCmd         *exec.Cmd
	incomingLogs     = make(chan string)
	anvilMux         sync.Mutex
)

func main() {

	http.HandleFunc("/", handleConnections)
	http.HandleFunc("/restart", handleAnvilRestart)

	go messageLoop()

	log.Println("Servidor iniciado na porta :8081")
	err := http.ListenAndServe(":8081", nil)

	if err != nil {
		fmt.Println("ListenAndServe: ", err)
		os.Exit(1)
	}
}
