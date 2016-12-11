package main

import "net/http"
import "golang.org/x/net/websocket"
import "log"
import "time"
import "strconv"
import "runtime"

func socketLive(ws *websocket.Conn) {
	var times int
	for {
		var msg string
		msg = "Hello" + strconv.Itoa(times)
		err := websocket.Message.Send(ws, msg)
		if err != nil {
			log.Println(err)
		}
		<-time.After(time.Second * 1)
		times++
	}
}

func main() {
	log.Println("Main", runtime.NumGoroutine())
	connectedCount := 0
	onConnected := func(ws *websocket.Conn) {
		// go func() {
		log.Println("Connected", runtime.NumGoroutine())
		err := websocket.Message.Send(ws, "test")
		if err != nil {
			log.Println(err)
		}
		connectedCount++
		socketLive(ws)
		log.Println(connectedCount, ws.Request().RemoteAddr)
		// }()
	}

	http.Handle("/", websocket.Handler(onConnected))
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Println(err)
	}
}
