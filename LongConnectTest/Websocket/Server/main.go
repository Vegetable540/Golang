package main

import "net/http"
import "golang.org/x/net/websocket"
import "log"

func socketLive(ws *websocket.Conn) {
	testCh := make(chan bool)
	for {
		test := <-testCh
		log.Println(test)
	}
}

func main() {
	connectedCount := 0
	onConnected := func(ws *websocket.Conn) {
		connectedCount++
		go socketLive(ws)
		log.Println(connectedCount, ws.LocalAddr().String()+ws.RemoteAddr().String())
	}

	http.Handle("/", websocket.Handler(onConnected))
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Println(err)
	}
}
