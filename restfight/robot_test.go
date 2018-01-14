package restfight

import (
	"testing"
)

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
