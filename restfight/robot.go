package restfight

import (
	"errors"
	"fmt"
	"math"
)

// CreateRobot based on given properties.
func CreateRobot(engineLevel int, shieldLevel int, weaponLevel int) (Robot, error) {

	var robot = Robot{
		Health:      10,
		MaxHealth:   10,
		Capacity:    0,
		MaxCapacity: 10,
		X:           0,
		Y:           0,
		MaxMoves:    0,
		Moves:       0,
		WeaponRange: 0,
		WeaponPower: 0,
		WeaponAmmo:  1,
	}

	// Parameter validation.
	if engineLevel < 0 || engineLevel > 2 {
		return robot, errors.New("INVALID_ENGINE_LEVEL")
	}
	if shieldLevel < 0 || shieldLevel > 2 {
		return robot, errors.New("INVALID_SHIELD_LEVEL")
	}
	if weaponLevel < 0 || weaponLevel > 2 {
		return robot, errors.New("INVALID_WEAPON_LEVEL")
	}

	// Ssssset!
	robot.WeaponLevel = weaponLevel
	robot.EngineLevel = engineLevel
	robot.ShieldLevel = shieldLevel

	// Setup engine.
	if engineLevel == 0 {
		robot.Capacity += 2
		robot.MaxMoves = 2
	}
	if engineLevel == 1 {
		robot.Capacity += 4
		robot.MaxMoves = 4
	}
	if engineLevel == 2 {
		robot.Capacity += 6
		robot.MaxMoves = 6
	}

	// Setup shield.
	if shieldLevel == 0 {
		robot.Capacity += 2
		robot.MaxHealth = 6
	}
	if shieldLevel == 1 {
		robot.Capacity += 4
		robot.MaxHealth = 10
	}
	if shieldLevel == 2 {
		robot.Capacity += 6
		robot.MaxHealth = 14
	}
	robot.Health = robot.MaxHealth

	// Setup weapon.
	if weaponLevel == 0 {
		robot.Capacity += 2
		robot.WeaponRange = 2
		robot.WeaponPower = 2
	}
	if weaponLevel == 1 {
		robot.Capacity += 4
		robot.WeaponRange = 4
		robot.WeaponPower = 4
	}
	if weaponLevel == 2 {
		robot.Capacity += 6
		robot.WeaponRange = 6
		robot.WeaponPower = 6
	}

	// Check capacity.
	if robot.Capacity > 10 {
		fmt.Printf("Robot capacity exceed: %d\n", robot.Capacity)
		return robot, errors.New("ROBOT_CAPACITY_EXCEED")
	}

	return robot, nil

}

// Scan surrounding arena.
func Scan() [ArenaSize][ArenaSize]Cell {

	return arena

}

// MoveRobot to a given position.
func MoveRobot(robotIndex int, x int, y int) (*Robot, error) {

	var robot *Robot

	if robotIndex >= cap(robots) {
		return robot, errors.New("ROBOT_INDEX_OUT_OF_BOUNDS")
	}

	if robotIndex >= len(robots) {
		return robot, errors.New("ROBOT_NOT_FOUND")
	}

	if status.Status != GameStatusRunning {
		return robot, errors.New("GAME_NOT_RUNNING")
	}

	if robotIndex != status.ActiveRobot {
		return robot, errors.New("NOT_YOUR_TURN")
	}

	robot = robots[robotIndex]

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
		return robot, errors.New("DIAGONAL_MOVES_NOT_ALLOWED")
	}

	// Only one step allowed.
	if math.Abs(float64(robot.X-x)) > 1 || math.Abs(float64(robot.Y-y)) > 1 {
		return robot, errors.New("ONLY_ONE_STEP_MOVES_ALLOWED")
	}

	// Avoid collision.
	if arena[x][y].Type == ArenaTypeRobot {
		return robot, errors.New("COLLISIONS_NOT_ALLOWED")
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

	if status.Status != GameStatusRunning {
		return errors.New("GAME_NOT_RUNNING")
	}

	if robotIndex != status.ActiveRobot {
		return errors.New("NOT_YOUR_TURN")
	}

	robot = robots[robotIndex]

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

	// Update game status.
	// TODO: From architecture perspective this is not the best location to update status.
	// This should happen outside this function.
	updateStatus()

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
