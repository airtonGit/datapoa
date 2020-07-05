package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write the file to the client.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the client.
	pongWait = 60 * time.Second

	// Send pings to client with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func reader(ws *websocket.Conn) {
	log.Println("WS Reader")
	defer ws.Close()
	ws.SetReadLimit(512)
	ws.SetReadDeadline(time.Now().Add(pongWait))
	ws.SetPongHandler(func(string) error { ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		log.Println("WS reader loop")
		_, _, err := ws.ReadMessage()
		if err != nil {
			log.Println("WS reader", err)
			break
		}
	}
}

func writer(ws *websocket.Conn) {
	log.Println("WS Writer")
	pingTicker := time.NewTicker(pingPeriod)
	defer func() {
		pingTicker.Stop()
		ws.Close()
	}()
	for {
		log.Println("WS loop")
		select {
		case <-pingTicker.C:
			log.Println("WS tick")
			ws.SetWriteDeadline(time.Now().Add(writeWait))
			if err := ws.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				log.Println("WS error", err)
				return
			}
		}
	}
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	log.Println("WS serveWs")
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WS ServeWs err", err)
		if _, ok := err.(websocket.HandshakeError); !ok {
			log.Println(err)
		}
		return
	}

	go writer(ws)
	reader(ws)
}
