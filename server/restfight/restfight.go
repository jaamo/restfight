/*
Core package for running the game.
*/
package restfight

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
)

// ArenaSize defines the size for the game arena. Arena is always square.
const ArenaSize = 10

// Arena
var arena [ArenaSize][ArenaSize]Cell

// Robots
var robots [2]Robot
var robotsInitialised = 0

// Status
var status Status

// NewGame starts new game.
func NewGame() {

	// Init empty arena.
	for x := 0; x < ArenaSize; x++ {
		for y := 0; y < ArenaSize; y++ {
			arena[x][y] = Cell{Type: ArenaTypeEmpty}
		}
	}

	// Init game status
	status = Status{
		GameID:            generateKey(1, 100),
		Status:            GameStatusWaitingForPlayers,
		ActiveRobot:       0,
		ActiveRobotStatus: ActiveRobotStatusWaiting,
	}

	robotsInitialised = 0

}

// JoinGame add a new robot to the arena and returns it. Return an error if game is full.
func JoinGame() (Robot, error) {

	var robot Robot

	if robotsInitialised == 2 {
		return robot, errors.New("GAME_FULL")
	}

	// Robot coordinates.
	x := 0
	y := 0
	if robotsInitialised == 1 {
		x = ArenaSize - 1
		y = ArenaSize - 1
	}

	// Create robot.
	robot = Robot{X: x, Y: y, RobotID: generateKey(robotsInitialised, 100)}
	robots[robotsInitialised] = robot

	// Add to arena.
	arena[x][y].Type = ArenaTypeRobot
	arena[x][y].Robot = &robots[robotsInitialised]

	robotsInitialised++

	return robot, nil

}

// MoveRobot moves a robot to give position.
func MoveRobot(robotIndex int, x int, y int) (Robot, error) {

	var robot Robot

	if robotIndex >= len(robots) {
		return robot, errors.New("ROBOT_INDEX_OUT_OF_BOUNDS")
	}

	// TÄÄ EI TOIMI!!
	if robots[robotIndex] == (Robot{}) {
		return robot, errors.New("ROBOT_NOT_FOUND")
	}

	robot = robots[robotIndex]
	// robot.X = 1

	return robot, nil

}

// GetStatus is only debugging atm.
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
