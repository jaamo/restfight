package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/jaamo/restfight/restfight"
)

type msg struct {
	Num int
}

// GameEvent represents a single event on the game (movement, shoot etc).
type GameEvent struct {
	EventType string          `json:"event_type"`
	Robot     restfight.Robot `json:"robot"`
}

// connections stores all active websocket connections
var connections []*websocket.Conn

// wsHandler handles WebSocket calls.
func wsHandler(w http.ResponseWriter, r *http.Request) {

	// Allow only connections from the same origin,.
	if r.Header.Get("Origin") != "http://"+r.Host {
		http.Error(w, "Origin not allowed", 403)
		return
	}

	// Init connection.
	conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
	}

	connections = append(connections, conn)

	// go echo(conn)

}

// broadcastSlice sends an event to all listening clients.
func broadcastEvent(gameEvent GameEvent) {

	fmt.Printf("Broadcast to clients!")

	for i := 0; i < len(connections); i++ {
		connections[i].WriteJSON(gameEvent)
	}

}

// broadcastStatus sends a game status to all listening clients.
func broadcastStatus() {

}

// echo function echoes all message from the network to all other connections
func echo(conn *websocket.Conn) {

	for {
		m := msg{}

		err := conn.ReadJSON(&m)
		if err != nil {
			fmt.Println("Error reading json.", err)
		}

		fmt.Printf("Got message: %#v\n", m)

		if err = conn.WriteJSON(m); err != nil {
			fmt.Println(err)
		}
	}

}
