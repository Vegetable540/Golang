package main

import "net/http"
import "golang.org/x/net/websocket"
import "log"

func socket(ws *websocket.Conn) {
	log.Print(ws.RemoteAddr)
}

func main() {
	http.ListenAndServe("9999", websocket.Handler(socket))
}
