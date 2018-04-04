package restfight

import (
	"errors"
	"math"
)

// Robots
var robots []Robot

// Robot object.
type Robot struct {
	RobotID     int `json:"robot_id"`
	Health      int `json:"health"`
	MaxHealth   int `json:"max_health"`
	Capacity    int `json:"capacity"`
	MaxCapacity int `json:"max_capacity"`
	X           int `json:"x"`
	Y           int `json:"y"`
	MaxMoves    int `json:"max_moves"`
	Moves       int `json:"moves"`
	WeaponRange int `json:"weapon_range"`
	WeaponPower int `json:"weapon_power"`
	WeaponAmmo  int `json:"weapon_ammo"`
	// RadarRange  int `json:"radar_range"`
}

// Radar object.
type Radar struct {
	Range int
}

// Scan returns arena.
func Scan() [ArenaSize][ArenaSize]Cell {

	return arena

}

// MoveRobot moves a robot to given position.
func MoveRobot(robotIndex int, x int, y int) (*Robot, error) {

	var robot *Robot

	if robotIndex >= cap(robots) {
		return robot, errors.New("ROBOT_INDEX_OUT_OF_BOUNDS")
	}

	if robotIndex >= len(robots) {
		return robot, errors.New("ROBOT_NOT_FOUND")
	}

	if robotIndex != turn {
		return robot, errors.New("NOT_YOUR_TURN")
	}

	robot = &robots[robotIndex]

	// Check moves left.
	if robot.Moves >= robot.MaxMoves {
		return robot, errors.New("OUT_OF_MOVES")
	}

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

	// Increase move counter.
	robot.Moves++

	// Move.
	forceMoveRobot(robot, x, y)

	return robot, nil

}

// GetRobotIndexByID return a robot by ID.
func GetRobotIndexByID(robotID int) (int, error) {

	for i := 0; i < len(robots); i++ {
		if robots[i].RobotID == robotID {
			return i, nil
		}
	}

	return 0, errors.New("ROBOT_NOT_FOUND")
}

// Shoot fires the weapon and shoots to given position.
func Shoot(robotIndex int, x int, y int) error {

	var robot *Robot

	if robotIndex >= cap(robots) {
		return errors.New("ROBOT_INDEX_OUT_OF_BOUNDS")
	}

	if robotIndex >= len(robots) {
		return errors.New("ROBOT_NOT_FOUND")
	}

	if robotIndex != turn {
		return errors.New("NOT_YOUR_TURN")
	}

	robot = &robots[robotIndex]

	if robot.WeaponAmmo <= 0 {
		return errors.New("OUT_OF_AMMO")
	}

	if x < 0 || x >= ArenaSize || y < 0 || y >= ArenaSize {
		return errors.New("OUT_OF_BOUNDS")
	}

	if math.Abs(float64(x-robot.X)) > float64(robot.WeaponRange) {
		return errors.New("OUT_OF_RANGE")
	}

	// Reduce ammo.
	robot.WeaponAmmo = 0

	// Get cell.
	cell := arena[x][y]

	// If robot found, reduce health.
	if cell.Type == ArenaTypeRobot {
		cell.Robot.Health -= robot.WeaponPower
	}

	return nil
}

// forceMoveRobot move robot to given position. Does not do any checks. Does not increase move count. Internal use only.
func forceMoveRobot(robot *Robot, x int, y int) {

	arena[robot.X][robot.Y].Type = ArenaTypeEmpty
	robot.X = x
	robot.Y = y
	arena[robot.X][robot.Y].Type = ArenaTypeRobot
	arena[robot.X][robot.Y].Robot = robot

}
