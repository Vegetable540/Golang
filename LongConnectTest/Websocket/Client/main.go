package main

import (
	"log"

	"time"

	"golang.org/x/net/websocket"
)

const ServerIP string = "192.168.184.131" //"192.168.222.128"

func ConnectServer() bool {
	ws, err := websocket.Dial("ws://"+ServerIP+":9999", "", "http://"+ServerIP+":9999")
	if err != nil {
		log.Println(err)
		return false
	}

	socketLive := func() {
		for {
			var msg string
			websocket.Message.Receive(ws, &msg)
			log.Println(msg)
			<-time.After(time.Second * 1)
		}
	}

	go socketLive()
	return true
}

func main() {
	for i := 0; i < 2000; i++ {
		ConnectServer()
	}

	blockCh := make(chan bool)
	block := <-blockCh
	log.Println(block)
}
