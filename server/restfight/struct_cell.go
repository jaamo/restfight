package restfight

// Cell represents a single cell in the arena.
type Cell struct {
	Type  int    `json:"type,omitempty"`
	Robot *Robot `json:"robot,omitempty"`
}

// ArenaTypeEmpty is constant for empty cell.
const ArenaTypeEmpty = 0

// ArenaTypeRobot is constant for a cell with a robot.
const ArenaTypeRobot = 1
