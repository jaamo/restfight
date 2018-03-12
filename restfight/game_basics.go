package restfight

import (
	"errors"
	"fmt"
	"strconv"
)

// Cell represents a single cell in the arena.
type Cell struct {
	Type  int    `json:"type"`
	X     int    `json:"x"`
	Y     int    `json:"y"`
	Robot *Robot `json:"robot"`
}

// ArenaTypeEmpty is constant for empty cell.
const ArenaTypeEmpty = 0

// ArenaTypeRobot is constant for a cell with a robot.
const ArenaTypeRobot = 1

// Status
var status Status

// ArenaSize defines the size for the game arena. Arena is always square.
const ArenaSize = 10

// Arena
var arena [ArenaSize][ArenaSize]Cell

/**
 * Status data model for the game.
 */
type Status struct {

	// Unique game id.
	GameID int `json:"game_id,omitempty"`

	// Game status: 0 = waiting for players, 1 = robot deployment, 2 = game is on, 3 = game over
	Status GameStatus `json:"status,omitempty"`

	// Active robot. 0 or 1.
	ActiveRobot int `json:"active_robot,omitempty"`

	// Active robot status. 0 = waiting, 1 = turn started
	ActiveRobotStatus ActiveRobotStatus `json:"active_robot_status,omitempty"`
}

/**
 * Datatype defining game status.
 */
type GameStatus int

/**
 * Game status constannts.
 */
const (
	GameStatusWaitingForPlayers GameStatus = iota
	GameStatusDeployment
	GameStatusRunning
	GameStatusGameOver
)

/**
 * Labels for game statuses.
 */
var GameStatusLabels = [...]string{
	"WAITING_FOR_PLAYERS",
	"DEPLOYMENT",
	"RUNNING",
	"GAME_OVER",
}

/**
 * Active robot status.
 */
type ActiveRobotStatus int

/**
 * Game status constannts.
 */
const (
	ActiveRobotStatusWaiting ActiveRobotStatus = iota
	ActiveRobotStatusPlaying
)

/**
 * Labels for game statuses.
 */
var ActiveRobotStatusLabels = [...]string{
	"WAITING",
	"PLAYING",
}

// NewGame starts new game.
func NewGame() {

	// Init empty arena.
	for x := 0; x < ArenaSize; x++ {
		for y := 0; y < ArenaSize; y++ {
			arena[x][y] = Cell{Type: ArenaTypeEmpty, X: x, Y: y}
		}
	}

	// Init game status
	status = Status{
		GameID:            generateKey(1, 100),
		Status:            GameStatusWaitingForPlayers,
		ActiveRobot:       0,
		ActiveRobotStatus: ActiveRobotStatusWaiting,
	}

	// Clear all players.
	robots = robots[:0]

}

// JoinGame add a new robot to the arena and returns it. Return an error if game is full.
func JoinGame() (Robot, error) {

	var robot Robot

	if len(robots) == 2 {
		return robot, errors.New("GAME_FULL")
	}

	// Robot coordinates.
	x := 0
	y := 0
	if len(robots) == 1 {
		x = ArenaSize - 1
		y = ArenaSize - 1
	}

	// Create robot.
	robot = Robot{X: x, Y: y, RobotID: generateKey(len(robots), 100), Radar: Radar{Range: 3}}
	robots = append(robots, robot)

	// Add to arena.
	arena[x][y].Type = ArenaTypeRobot
	arena[x][y].Robot = &robot

	return robot, nil

}

// GetStatus is only debugging atm.
func GetStatus() Status {

	//var arena [ArenaSize][ArenaSize]Cell

	var output string
	for x := 0; x < ArenaSize; x++ {
		for y := 0; y < ArenaSize; y++ {
			output += strconv.Itoa(arena[x][y].Type) + " "
		}
		output += "\n"
	}
	fmt.Println(output)

	return status
}
