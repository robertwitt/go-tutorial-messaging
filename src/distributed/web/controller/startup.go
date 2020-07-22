package controller

import (
	"net/http"
)

var ws = newWebsocketController()

// Initialize initialize
func Initialize() {
	registerRoutes()
	registerFileServers()
}

func registerRoutes() {
	http.HandleFunc("/ws", ws.handleMessage)
}

func registerFileServers() {
	// TODO
}
