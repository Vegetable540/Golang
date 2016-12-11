package main

import (
	"golang.org/x/net/websocket"
	"log"
)

const ServerIP string = "127.0.0.1" //"192.168.222.128"

func ConnectServer() bool {
	ws, err := websocket.Dial("ws://"+ServerIP+":9999", "", "http://"+ServerIP+":9999")
	if err != nil {
		log.Println(err)
		return false
	}

	socketLive := func() {
		log.Println(ws.LocalAddr().String() + ws.RemoteAddr().String())
		test := make(chan bool)
		for {
			t := <-test
			log.Println(t)
		}
	}

	go socketLive()
	return true
}

func main() {
	for {
		if !ConnectServer() {
			return
		}
	}
}
