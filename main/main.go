package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jaamo/restfight/restfight"

	"github.com/gorilla/mux"
)

// gameAPIError represents API error
type gameAPIError struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

// apiGetStatus return complete game status.
func apiGetStatus(w http.ResponseWriter, r *http.Request) {
	broadcastEvent(GameEvent{EventType: "STATUS_QUERIED"})
	json.NewEncoder(w).Encode(restfight.GetStatus())
}

// apiJoinGame registers a new player.
func apiJoinGame(w http.ResponseWriter, r *http.Request) {
	robot, error := restfight.JoinGame()
	if error != nil {
		apiError(w, error.Error(), "")
	} else {
		broadcastEvent(GameEvent{EventType: "JOIN_GAME", Robot: robot})
		json.NewEncoder(w).Encode(robot)
	}
}

// apiEchoDebug echoes a message to clients
func apiEchoDebug(w http.ResponseWriter, r *http.Request) {
	broadcastEvent(GameEvent{EventType: "DEBUG"})
}

// apiError is a helper function to return REST error message.
func apiError(w http.ResponseWriter, error string, message string) {
	json.NewEncoder(w).Encode(gameAPIError{Error: error, Message: message})
}

// main
func main() {

	fmt.Println("Starting server on port 8000.")

	// Start new game.
	restfight.NewGame()

	// Create router,
	router := mux.NewRouter()

	// Register REST routes.
	router.HandleFunc("/echodebug", apiEchoDebug).Methods("GET")
	router.HandleFunc("/join", apiJoinGame).Methods("GET")
	router.HandleFunc("/status", apiGetStatus).Methods("GET")

	// Setup static file serving.
	s := http.StripPrefix("/viewer/", http.FileServer(http.Dir("./viewer/")))
	router.PathPrefix("/viewer/").Handler(s)
	http.Handle("/", router)

	// Handle WebSocket connections.
	router.HandleFunc("/socket", wsHandler)

	// Start the server.
	log.Fatal(http.ListenAndServe(":8000", router))

}
