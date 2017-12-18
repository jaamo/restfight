package restfight

// Status struct.
type Status struct {
	GameID string `json:"game_id,omitempty"`
	Turn   string `json:"turn,omitempty"`
}

func GetStatus() int {
	return 2
}
