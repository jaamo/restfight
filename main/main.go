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

	robotIndex := -1

	robotIDParam, ok := r.URL.Query()["robot_id"]
	if ok || len(robotIDParam) > 0 {
		robotID, _ := strconv.Atoi(robotIDParam[0])
		robotIndex, _ = restfight.GetRobotIndexByID(robotID)
	}

	broadcastEvent(GameEvent{EventType: "STATUS", Status: restfight.GetStatus(-1)})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(restfight.GetStatus(robotIndex))

}

// apiJoinGame registers a new player.
func apiJoinGame(w http.ResponseWriter, r *http.Request) {

	engineLevelParam, ok := r.URL.Query()["engineLevel"]
	if !ok || len(engineLevelParam) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(GameError{Error: "PARAMETER_MISSING", Message: "Parameter engineLevel missing."})
		return
	}

	shieldLevelParam, ok := r.URL.Query()["shieldLevel"]
	if !ok || len(shieldLevelParam) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(GameError{Error: "PARAMETER_MISSING", Message: "Parameter shieldLevel missing."})
		return
	}

	weaponLevelParam, ok := r.URL.Query()["weaponLevel"]
	if !ok || len(weaponLevelParam) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(GameError{Error: "PARAMETER_MISSING", Message: "Parameter weaponLevel missing."})
		return
	}

	engineLevel, _ := strconv.Atoi(engineLevelParam[0])
	shieldLevel, _ := strconv.Atoi(shieldLevelParam[0])
	weaponLevel, _ := strconv.Atoi(weaponLevelParam[0])

	robot, error := restfight.JoinGame(engineLevel, shieldLevel, weaponLevel)

	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
		apiError(w, error.Error(), "")
	} else {
		broadcastEvent(GameEvent{EventType: "JOIN_GAME", Robot: robot})
		broadcastEvent(GameEvent{EventType: "STATUS", Status: restfight.GetStatus(robot.RobotIndex)})
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
	json.NewEncoder(w).Encode(restfight.GetStatus(-1))
	broadcastEvent(GameEvent{EventType: "NEW_GAME"})
}

// apiMove moves robot new position
func apiMove(w http.ResponseWriter, r *http.Request) {

	xParams, ok := r.URL.Query()["x"]
	if !ok || len(xParams) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(GameError{Error: "PARAMETER_MISSING", Message: "Parameter x missing."})
		return
	}

	yParams, ok := r.URL.Query()["y"]
	if !ok || len(yParams) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(GameError{Error: "PARAMETER_MISSING", Message: "Parameter y missing."})
		return
	}

	robotIDParam, ok := r.URL.Query()["robot_id"]
	if !ok || len(robotIDParam) < 1 {
		w.WriteHeader(http.StatusBadRequest)
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
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(GameError{Error: err.Error(), Message: fmt.Sprintf("You tried move robot %d (index %d) to  %d x %d.", robotID, robotIndex, x, y)})
		return
	}

	json.NewEncoder(w).Encode(robot)
	broadcastEvent(GameEvent{EventType: "ROBOT_MOVED", Robot: *robot})

}

func apiShoot(w http.ResponseWriter, r *http.Request) {

	xParams, ok := r.URL.Query()["x"]
	if !ok || len(xParams) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(GameError{Error: "PARAMETER_MISSING", Message: "Parameter x missing."})
		return
	}

	yParams, ok := r.URL.Query()["y"]
	if !ok || len(yParams) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(GameError{Error: "PARAMETER_MISSING", Message: "Parameter y missing."})
		return
	}

	robotIDParam, ok := r.URL.Query()["robot_id"]
	if !ok || len(robotIDParam) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(GameError{Error: "PARAMETER_MISSING", Message: "Parameter robot_id missing."})
		return
	}

	x, _ := strconv.Atoi(xParams[0])
	y, _ := strconv.Atoi(yParams[0])
	robotID, _ := strconv.Atoi(robotIDParam[0])
	var robotIndex, _ = restfight.GetRobotIndexByID(robotID)

	// Shoot.
	err := restfight.Shoot(robotIndex, x, y)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(GameError{Error: err.Error(), Message: fmt.Sprintf("You tried to shoot to %d x %d.", x, y)})
		return
	}

	broadcastEvent(GameEvent{EventType: "SHOOT", X: x, Y: y})
	broadcastEvent(GameEvent{EventType: "STATUS", Status: restfight.GetStatus(robotIndex)})
	json.NewEncoder(w).Encode(restfight.GetStatus(robotIndex))

}

// apiMove moves robot new position
func apiEndTurn(w http.ResponseWriter, r *http.Request) {

	robotIDParam, ok := r.URL.Query()["robot_id"]
	if !ok || len(robotIDParam) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(GameError{Error: "PARAMETER_MISSING", Message: "Parameter robot_id missing."})
		return
	}

	robotID, _ := strconv.Atoi(robotIDParam[0])
	var robotIndex, _ = restfight.GetRobotIndexByID(robotID)

	if restfight.CanPlay(robotIndex) {
		restfight.ToggleTurn()
		broadcastEvent(GameEvent{EventType: "NEW_TURN", Status: restfight.GetStatus(robotIndex)})
	} else {
		w.WriteHeader(http.StatusBadRequest)
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
	router.HandleFunc("/shoot", apiShoot).Methods("GET")

	// Setup static file serving.
	s := http.StripPrefix("/viewer/", http.FileServer(http.Dir("./viewer/")))
	router.PathPrefix("/viewer/").Handler(s)
	http.Handle("/", router)

	// Handle WebSocket connections.
	router.HandleFunc("/socket", wsHandler)

	// Start the server.
	log.Fatal(http.ListenAndServe(":8000", router))

}
