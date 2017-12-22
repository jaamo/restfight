package restfight

// Status struct.
type Status struct {
	GameID string `json:"game_id,omitempty"`
	Turn   string `json:"turn,omitempty"`
}

/**
 * Size of the arena.
 */
const ArenaSize = 100

/**
 * Arena type. Empty.
 */
const ArenaTypeEmpty = 0

/**
 * Single cell object.
 */
type Cell struct {
	Type  int8   `json:"type,omitempty"`
	Robot string `json:"robot,omitempty"`
}

/**
 * Get game status.
 */
func GetStatus() int {

	var arena [ArenaSize][ArenaSize]Cell

	for x := 0; x < ArenaSize; x++ {
		for y := 0; y < ArenaSize; y++ {
			arena[x][y] = Cell{Type: ArenaTypeEmpty}
		}
	}

	return 2
}
