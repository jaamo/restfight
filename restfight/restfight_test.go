package restfight

import (
	"testing"
)

func TestNewGameArenaSize(t *testing.T) {

	NewGame()

	if len(arena) != ArenaSize {
		t.Errorf("Arena size is incorrect %d != %d", len(arena), ArenaSize)
	}

	if len(arena[0]) != ArenaSize {
		t.Errorf("Arena size is incorrect %d != %d", len(arena[0]), ArenaSize)
	}

}
func TesJoinGameHasRobots(t *testing.T) {

	NewGame()
	JoinGame()
	JoinGame()

	numberOfRobots := 0
	for x := 0; x < ArenaSize; x++ {
		for y := 0; y < ArenaSize; y++ {
			if arena[x][y].Type == ArenaTypeRobot {
				numberOfRobots++
			}
		}
	}

	if numberOfRobots != 2 {
		t.Errorf("Incorrect amount of robots. Found %d, expected %d.", numberOfRobots, 2)
	}

}
func TestNewGameRobotsAreOk(t *testing.T) {

	NewGame()
	JoinGame()
	JoinGame()

	if robots[0].X != 0 || robots[0].Y != 0 {
		t.Errorf("Robot 1 position is incorrect. Found %d x %d, expected 0 x 0.", robots[0].X, robots[0].Y)
	}

	if robots[1].X != ArenaSize-1 || robots[1].Y != ArenaSize-1 {
		t.Errorf("Robot 2 position is incorrect. Found %d x %d, expected %d x %d.", robots[1].X, robots[1].Y, ArenaSize-1, ArenaSize-1)
	}

}

func TestNewGameLimitRobots(t *testing.T) {

	NewGame()

	_, error := JoinGame()
	if error != nil {
		t.Errorf("Tried to add the first robot but got an error: %s.", error)
	}

	_, error = JoinGame()
	if error != nil {
		t.Errorf("Tried to add a second robot but got an error: %s.", error)
	}

	_, error = JoinGame()
	if error == nil {
		t.Errorf("Game was supposed to be full but it wasn't.")
	}

}

func TestMoveRobot(t *testing.T) {

	NewGame()

	// Shoudn't be any robots at this point.
	_, err := MoveRobot(0, 0, 0)
	if err == nil || err.Error() != "ROBOT_NOT_FOUND" {
		t.Errorf("Was expecting error ROBOT_NOT_FOUND on index 0, got %s", err)
	}

	_, err = MoveRobot(1, 0, 0)
	if err == nil || err.Error() != "ROBOT_NOT_FOUND" {
		t.Errorf("Was expecting error ROBOT_NOT_FOUND, got %s", err)
	}
	_, err = MoveRobot(2, 0, 0)
	if err == nil || err.Error() != "ROBOT_INDEX_OUT_OF_BOUNDS" {
		t.Errorf("Was expecting error ROBOT_INDEX_OUT_OF_BOUNDS, got %s", err)
	}

	// Add robot. Should be only one.
	JoinGame()
	_, err = MoveRobot(1, 0, 0)
	if err == nil || err.Error() != "ROBOT_NOT_FOUND" {
		t.Errorf("Was expecting error ROBOT_NOT_FOUND, got %s", err)
	}

	JoinGame()

	// Try some invalid moves.
	_, err = MoveRobot(0, 2, 0)
	if err == nil || err.Error() != "INVALID_MOVE" {
		t.Errorf("Was expecting error INVALID_MOVE, got %s", err)
	}
	_, err = MoveRobot(0, 0, 5)
	if err == nil || err.Error() != "INVALID_MOVE" {
		t.Errorf("Was expecting error INVALID_MOVE, got %s", err)
	}
	_, err = MoveRobot(0, 8, 5)
	if err == nil || err.Error() != "INVALID_MOVE" {
		t.Errorf("Was expecting error INVALID_MOVE, got %s", err)
	}
	_, err = MoveRobot(0, 1, 1)
	if err == nil || err.Error() != "INVALID_MOVE" {
		t.Errorf("Was expecting error INVALID_MOVE, got %s", err)
	}

	// Out of bounds.
	_, err = MoveRobot(0, -1, 0)
	if err == nil || err.Error() != "OUT_OF_BOUNDS" {
		t.Errorf("Was expecting error OUT_OF_BOUNDS, got %s", err)
	}
	_, err = MoveRobot(0, 0, -1)
	if err == nil || err.Error() != "OUT_OF_BOUNDS" {
		t.Errorf("Was expecting error OUT_OF_BOUNDS, got %s", err)
	}
	_, err = MoveRobot(0, -1, -1)
	if err == nil || err.Error() != "OUT_OF_BOUNDS" {
		t.Errorf("Was expecting error OUT_OF_BOUNDS, got %s", err)
	}

	// Valid moves
	var robot *Robot
	robot, err = MoveRobot(0, 0, 1)
	if robot.X != 0 || robot.Y != 1 {
		t.Errorf("Move failed. Expected position %d x %d, got %d x %d. Error %s", 0, 1, robot.X, robot.Y, err)
	}

	// Collision
	forceMoveRobot(&robots[0], 0, 0)
	forceMoveRobot(&robots[1], 0, 1)
	GetStatus()
	_, err = MoveRobot(0, 0, 1)
	if err == nil || err.Error() != "INVALID_MOVE" {
		t.Errorf("Was expecting error INVALID_MOVE, got %s", err)
	}

}
func TestScan(t *testing.T) {

	NewGame()
	JoinGame()
	JoinGame()

	// Do scan on the top left corner.
	s, _ := Scan(robots[0])

	if len(s) != 4 {
		t.Errorf("Too many scan results: %d. Was expecting 4.", len(s))
	}

	if s[0].Type != ArenaTypeRobot {
		t.Errorf("Expected robot on 0 x 0. Got %d.", s[0].Type)
	}

	// Move another robot to range.
	forceMoveRobot(&robots[1], 0, 1)
	s, _ = Scan(robots[0])

	if s[1].Type != ArenaTypeRobot {
		t.Errorf("Expected robot on 1 x 0. Got %d.", s[1].Type)
	}

}
func TestKeyGeneration(t *testing.T) {
	len := 100
	key := generateKey(1, len)
	if key < len {
		t.Errorf("Generated key %d is not what expected. Key was expected to start with 1 and be between 100 - 199", key)
	}
	if key > 199 {
		t.Errorf("Generated key %d is not what expected. Key was expected to start with 1 and be between 100 - 199", key)
	}

	key = generateKey(2, len)
	if key < 2*len {
		t.Errorf("Generated key %d is not what expected. Key was expected to start with 1 and be between 200 - 299", key)
	}
	if key > 299 {
		t.Errorf("Generated key %d is not what expected. Key was expected to start with 1 and be between 200 - 299", key)
	}

}
func TestStatus(t *testing.T) {

	var status = GetStatus()
	if status != 2 {
		t.Errorf("Error %d != %d", status, 2)
	}

}
