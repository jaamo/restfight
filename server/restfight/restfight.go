/*
Core package for running the game.
*/
package restfight

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"strconv"
)

// ArenaSize defines the size for the game arena. Arena is always square.
const ArenaSize = 10

// Arena
var arena [ArenaSize][ArenaSize]Cell

// Robots
var robots []Robot

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
	robot = Robot{X: x, Y: y, RobotID: generateKey(len(robots), 100)}
	robots = append(robots, robot)

	// Add to arena.
	arena[x][y].Type = ArenaTypeRobot
	arena[x][y].Robot = &robot

	return robot, nil

}

// MoveRobot moves a robot to give position.
func MoveRobot(robotIndex int, x int, y int) (*Robot, error) {

	var robot *Robot

	if robotIndex >= cap(robots) {
		return robot, errors.New("ROBOT_INDEX_OUT_OF_BOUNDS")
	}

	if robotIndex >= len(robots) {
		return robot, errors.New("ROBOT_NOT_FOUND")
	}

	robot = &robots[robotIndex]

	// Check out of bounds.
	if x < 0 || x >= ArenaSize || y < 0 || y >= ArenaSize {
		return robot, errors.New("OUT_OF_BOUNDS")
	}

	// Diagonal move not allowed.
	if robot.X != x && robot.Y != y {
		return robot, errors.New("INVALID_MOVE")
	}

	// Only one step allowed.
	if math.Abs(float64(robot.X-x)) > 1 || math.Abs(float64(robot.Y-y)) > 1 {
		return robot, errors.New("INVALID_MOVE")
	}

	// Avoid collision.
	if arena[x][y].Type == ArenaTypeRobot {
		return robot, errors.New("INVALID_MOVE")
	}

	forceMoveRobot(robot, x, y)

	return robot, nil

}

// forceMoveRobot move robot to given position. Does not do any checks.
func forceMoveRobot(robot *Robot, x int, y int) {

	arena[robot.X][robot.Y].Type = ArenaTypeEmpty
	robot.X = x
	robot.Y = y
	arena[robot.X][robot.Y].Type = ArenaTypeRobot
	arena[robot.X][robot.Y].Robot = robot

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
