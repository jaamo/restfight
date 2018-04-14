package restfight

import (
	"testing"
)

func TestGetRobotIndexByID(t *testing.T) {

	NewGame()

	var robot, _ = JoinGame(1, 1, 0)
	robotIndex, _ := GetRobotIndexByID(robot.RobotID)
	if robotIndex != 0 {
		t.Errorf("The first robot should have index 0. Got %d.", robotIndex)
	}

	var robot2, _ = JoinGame(1, 1, 0)
	robotIndex2, _ := GetRobotIndexByID(robot2.RobotID)
	if robotIndex2 != 1 {
		t.Errorf("The second robot should have index 1. Got %d.", robotIndex)
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
	JoinGame(1, 1, 0)
	_, err = MoveRobot(1, 0, 0)
	if err == nil || err.Error() != "ROBOT_NOT_FOUND" {
		t.Errorf("Was expecting error ROBOT_NOT_FOUND, got %s", err)
	}

	// Turn should still be -1.
	if status.ActiveRobot != 0 && status.Status == GameStatusWaitingForPlayers {
		t.Errorf("Turn should be -1, was %d", status.ActiveRobot)
	}

	JoinGame(1, 1, 0)

	// Turn should be 0.
	if status.ActiveRobot != 0 {
		t.Errorf("Turn should be 0, was %d", status.ActiveRobot)
	}

	// Turn check.
	_, err = MoveRobot(1, 0, 0)
	if err == nil || err.Error() != "NOT_YOUR_TURN" {
		t.Errorf("Was expecting error NOT_YOUR_TURN, got %s", err)
	}

	// Try some invalid moves.
	_, err = MoveRobot(0, 2, 0)
	if err == nil || err.Error() != "ONLY_ONE_STEP_MOVES_ALLOWED" {
		t.Errorf("Was expecting error ONLY_ONE_STEP_MOVES_ALLOWED, got %s", err)
	}
	_, err = MoveRobot(0, 0, 5)
	if err == nil || err.Error() != "ONLY_ONE_STEP_MOVES_ALLOWED" {
		t.Errorf("Was expecting error ONLY_ONE_STEP_MOVES_ALLOWED, got %s", err)
	}
	_, err = MoveRobot(0, 8, 5)
	if err == nil || err.Error() != "DIAGONAL_MOVES_NOT_ALLOWED" {
		t.Errorf("Was expecting error DIAGONAL_MOVES_NOT_ALLOWED, got %s", err)
	}
	_, err = MoveRobot(0, 1, 1)
	if err == nil || err.Error() != "DIAGONAL_MOVES_NOT_ALLOWED" {
		t.Errorf("Was expecting error DIAGONAL_MOVES_NOT_ALLOWED, got %s", err)
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

	// Two more moves.
	robot, err = MoveRobot(0, 0, 2)
	robot, err = MoveRobot(0, 0, 3)
	robot, err = MoveRobot(0, 0, 4)

	// Should be out of moves now.
	robot, err = MoveRobot(0, 0, 5)
	if err == nil || err.Error() != "OUT_OF_MOVES" {
		t.Errorf("Was expecting error OUT_OF_MOVES, got %s", err)
	}

	// New turn. Should be able to move again.
	ToggleTurn()
	ToggleTurn()
	robot, err = MoveRobot(0, 0, 5)
	if err != nil {
		t.Errorf("Wasn't expecting errors, got %s", err)
	}

	// Collision
	robot.Moves = 0
	forceMoveRobot(robots[0], 0, 0)
	forceMoveRobot(robots[1], 0, 1)
	// GetStatus()
	_, err = MoveRobot(0, 0, 1)
	if err == nil || err.Error() != "COLLISIONS_NOT_ALLOWED" {
		t.Errorf("Was expecting error COLLISIONS_NOT_ALLOWED, got %s", err)
	}

}

func TestShoot(t *testing.T) {

	NewGame()
	JoinGame(1, 1, 0)
	JoinGame(1, 1, 0)

	var err error

	// Player 1 can't shoot, not his turn.
	err = Shoot(1, 1, 1)
	if err == nil || err.Error() != "NOT_YOUR_TURN" {
		t.Errorf("Was expecting error NOT_YOUR_TURN, got %s", err)
	}

	// Shoot outside arena.
	err = Shoot(0, -1, 0)
	if err == nil || err.Error() != "OUT_OF_BOUNDS" {
		t.Errorf("Was expecting error OUT_OF_BOUNDS, got %s", err)
	}

	err = Shoot(0, 100, 0)
	if err == nil || err.Error() != "OUT_OF_BOUNDS" {
		t.Errorf("Was expecting error OUT_OF_BOUNDS, got %s", err)
	}

	err = Shoot(0, 0, -1)
	if err == nil || err.Error() != "OUT_OF_BOUNDS" {
		t.Errorf("Was expecting error OUT_OF_BOUNDS, got %s", err)
	}

	err = Shoot(0, 0, 1000)
	if err == nil || err.Error() != "OUT_OF_BOUNDS" {
		t.Errorf("Was expecting error OUT_OF_BOUNDS, got %s", err)
	}

	forceMoveRobot(robots[1], 2, 2)
	// GetStatus(-1)

	// Should run out of ammo.
	err = Shoot(0, 1, 1)
	err = Shoot(0, 1, 1)
	if err == nil || err.Error() != "OUT_OF_AMMO" {
		t.Errorf("Was expecting error OUT_OF_AMMO, got %s", err)
	}
	robots[0].WeaponAmmo = 1

	err = Shoot(0, 2, 2)
	if err != nil {
		t.Errorf("Did not expect an error on succesful shoot.")
	}
	if robots[1].Health != robots[1].MaxHealth-robots[0].WeaponPower {
		t.Errorf("Health was supposed to be %d but it was %d.", robots[1].MaxHealth-robots[0].WeaponPower, robots[1].Health)
	}

}
