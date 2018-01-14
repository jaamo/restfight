package restfight

import (
	"errors"
	"math"
)

// Robots
var robots []Robot

// Robot object.
type Robot struct {
	RobotID     int   `json:"robot_id,omitempty"`
	Health      int   `json:"health,omitempty"`
	MaxHealth   int   `json:"max_health,omitempty"`
	Capacity    int   `json:"capacity,omitempty"`
	MaxCapacity int   `json:"max_capacity,omitempty"`
	X           int   `json:"x,omitempty"`
	Y           int   `json:"y,omitempty"`
	Radar       Radar `json:"radar,omitempty"`
}

// Radar object.
type Radar struct {
	Range int
}

// Scan returns an area from arena. Uses radar dimension from given robot.
func Scan(robot Robot) ([]Cell, error) {

	var cells []Cell
	for x := 0; x < robot.Radar.Range; x++ {
		for y := 0; y < robot.Radar.Range; y++ {
			// fmt.Printf("%d x %d\n", x, y)
			xOffset := x + robot.X - robot.Radar.Range/2
			yOffset := y + robot.Y - robot.Radar.Range/2
			if xOffset >= 0 && yOffset >= 0 && xOffset < ArenaSize && yOffset < ArenaSize {
				cells = append(cells, arena[xOffset][yOffset])
			}
		}
	}

	return cells, nil

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
