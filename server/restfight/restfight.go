package restfight

import (
	"fmt"
	"math/rand"
	"strconv"
)

// Status struct.
type Status struct {

	// Unique game id.
	GameID int `json:"game_id,omitempty"`

	// Game status: 0 = waiting for players, 1 = robot deployment, 2 = game is on, 3 = game over
	Status int `json:"status,omitempty"`

	// Active robot. 0 or 1.
	ActiveRobot int `json:"active_robot,omitempty"`

	// Active robot status. 0 = waiting, 1 = turn started
	ActiveRobotStatus int `json:"active_robot_status,omitempty"`
}

/**
 * Size of the arena.
 */
const ArenaSize = 10

/**
 * Arena type. Empty.
 */
const ArenaTypeEmpty = 0
const ArenaTypeRobot = 1

/**
 * Single cell object.
 */
type Cell struct {
	Type  int    `json:"type,omitempty"`
	Robot *Robot `json:"robot,omitempty"`
}

/**
 * Single cell object.
 */
type Robot struct {
	RobotID     int `json:"robot_id,omitempty"`
	Health      int `json:"health,omitempty"`
	MaxHealth   int `json:"max_health,omitempty"`
	Capacity    int `json:"capacity,omitempty"`
	MaxCapacity int `json:"max_capacity,omitempty"`
	X           int `json:"x,omitempty"`
	Y           int `json:"y,omitempty"`
}

var arena [ArenaSize][ArenaSize]Cell
var robots [2]Robot
var status Status

/**
 * Create new game.
 */
func NewGame() {

	// Init empty arena.
	for x := 0; x < ArenaSize; x++ {
		for y := 0; y < ArenaSize; y++ {
			arena[x][y] = Cell{Type: ArenaTypeEmpty}
		}
	}

	// Add robots.
	for i := 0; i < len(robots); i++ {
		x := 0
		y := 0
		if i == 1 {
			x = ArenaSize - 1
			y = ArenaSize - 1
		}
		robots[i] = Robot{X: x, Y: y, RobotID: generateKey(i, 100)}
		arena[robots[i].X][robots[i].Y].Type = ArenaTypeRobot
		arena[robots[i].X][robots[i].Y].Robot = &robots[i]
	}

	// Init game status
	status = Status{GameID: generateKey, Turn: 0}

}

/**
 * Get game status.
 */
func GetStatus() int {

	//var arena [ArenaSize][ArenaSize]Cell

	var output string
	for x := 0; x < ArenaSize; x++ {
		for y := 0; y < ArenaSize; y++ {
			output += strconv.Itoa(arena[x][y].Type) + " "
		}
		output += "\n"
	}
	fmt.Println(output)

	return 2
}

/**
 * Generate key.
 */
func generateKey(base int, length int) int {
	return base*length + rand.Intn(length/10)
}
