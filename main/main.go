package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"./restfight"

	"github.com/gorilla/mux"
)

// This implementation is based on this:
// https://www.codementor.io/codehakase/building-a-restful-api-with-golang-a6yivzqdo

// Status struct.
type Status struct {
	GameID string `json:"game_id,omitempty"`
	Turn   string `json:"turn,omitempty"`
}

/**
 * Return game status.
 */
func apiGetStatus(w http.ResponseWriter, r *http.Request) {
	var status = Status{GameID: "kala", Turn: "auto"}
	json.NewEncoder(w).Encode(status)
}

/**
 * Main function.
 */
func main() {

	fmt.Println("running")

	fmt.Println(restfight.GetStatus())

	// Create router,
	router := mux.NewRouter()

	// Register routes.
	router.HandleFunc("/status", apiGetStatus).Methods("GET")

	// Start the server.
	log.Fatal(http.ListenAndServe(":8000", router))

}
