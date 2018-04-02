package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

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

// apiEchoDebug resets the game
func apiNewGame(w http.ResponseWriter, r *http.Request) {
	restfight.NewGame()
	json.NewEncoder(w).Encode(restfight.GetStatus())
	broadcastEvent(GameEvent{EventType: "NEW_GAME"})
}

// apiMove moves robot new position
func apiMove(w http.ResponseWriter, r *http.Request) {

	xParams, ok := r.URL.Query()["x"]
	if !ok || len(xParams) < 1 {
		json.NewEncoder(w).Encode(GameError{Error: "PARAMETER_MISSING", Message: "Parameter x missing."})
		return
	}

	yParams, ok := r.URL.Query()["y"]
	if !ok || len(yParams) < 1 {
		json.NewEncoder(w).Encode(GameError{Error: "PARAMETER_MISSING", Message: "Parameter y missing."})
		return
	}

	robotIDParam, ok := r.URL.Query()["robot_id"]
	if !ok || len(robotIDParam) < 1 {
		json.NewEncoder(w).Encode(GameError{Error: "PARAMETER_MISSING", Message: "Parameter robot_id missing."})
		return
	}

	x, _ := strconv.Atoi(xParams[0])
	y, _ := strconv.Atoi(yParams[0])
	robotID, _ := strconv.Atoi(robotIDParam[0])
	var robotIndex, _ = restfight.GetRobotIndexByID(robotID)

	// Move robot.
	robot, err := restfight.MoveRobot(robotIndex, x, y)

	if err != nil {
		json.NewEncoder(w).Encode(GameError{Error: err.Error(), Message: fmt.Sprintf("You triedt move robot %d (index %d) to  %d x %d.", robotID, robotIndex, x, y)})
		return
	}

	json.NewEncoder(w).Encode(robot)
	broadcastEvent(GameEvent{EventType: "ROBOT_MOVED", Robot: *robot})

}

// apiMove moves robot new position
func apiEndTurn(w http.ResponseWriter, r *http.Request) {

	robotIDParam, ok := r.URL.Query()["robot_id"]
	if !ok || len(robotIDParam) < 1 {
		json.NewEncoder(w).Encode(GameError{Error: "PARAMETER_MISSING", Message: "Parameter robot_id missing."})
		return
	}

	robotID, _ := strconv.Atoi(robotIDParam[0])
	var robotIndex, _ = restfight.GetRobotIndexByID(robotID)

	if restfight.CanPlay(robotIndex) {
		restfight.ToggleTurn()
		broadcastEvent(GameEvent{EventType: "NEW_TURN"})
	} else {
		json.NewEncoder(w).Encode(GameError{Error: "NOT_YOUR_TURN", Message: "Not your turn."})
	}

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
	router.HandleFunc("/new", apiNewGame).Methods("GET")
	router.HandleFunc("/move", apiMove).Methods("GET")
	router.HandleFunc("/endturn", apiEndTurn).Methods("GET")

	// Setup static file serving.
	s := http.StripPrefix("/viewer/", http.FileServer(http.Dir("./viewer/")))
	router.PathPrefix("/viewer/").Handler(s)
	http.Handle("/", router)

	// Handle WebSocket connections.
	router.HandleFunc("/socket", wsHandler)

	// Start the server.
	log.Fatal(http.ListenAndServe(":8000", router))

}
