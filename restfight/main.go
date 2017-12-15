package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Status struct.
type Status struct {
	gameId string `json:"id,omitempty"`,
	turn   string `json:"turn,omitempty"`
}

/**
 * Return game status.
 */
func getStatus(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

/**
 * Main function.
 */
func main() {

	// Create router,
	router := mux.NewRouter()

	// Register routes.
	router.HandleFunc("/people", getStatus).Methods("GET")
	// router.HandleFunc("/people/{id}", fn).Methods("POST")

	// Start the server.
	log.Fatal(http.ListenAndServe(":8000", router))

}
