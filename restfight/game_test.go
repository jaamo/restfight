package restfight

import "testing"

func TestCanPlay(t *testing.T) {

	NewGame()
	JoinGame(1, 1, 0)
	JoinGame(1, 1, 0)

	if CanPlay(0) == false {
		t.Errorf("New game. Player 0 should be active.")
	}

	if CanPlay(1) == true {
		t.Errorf("New game. Player 1 shouldn't be active.")
	}

	ToggleTurn()

	if CanPlay(1) == false {
		t.Errorf("2nd turn. Player 1 should be active.")
	}

	if CanPlay(0) == true {
		t.Errorf("2nd turn. Player 0 shouldn't be active.")
	}

}
func TestToggleTurn(t *testing.T) {

	NewGame()
	JoinGame(1, 1, 0)
	JoinGame(1, 1, 0)

	if status.ActiveRobot != 0 {
		t.Errorf("New game. Turn should be 0.")
	}

	ToggleTurn()

	if status.ActiveRobot != 1 {
		t.Errorf("2nd turn. Turn should be 1.")
	}

	ToggleTurn()

	if status.ActiveRobot != 0 {
		t.Errorf("3rd turn. Turn should be 0.")
	}

}

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
	JoinGame(1, 1, 0)
	JoinGame(1, 1, 0)

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
	JoinGame(1, 1, 0)
	JoinGame(1, 1, 0)

	if robots[0].X != 0 || robots[0].Y != 0 {
		t.Errorf("Robot 1 position is incorrect. Found %d x %d, expected 0 x 0.", robots[0].X, robots[0].Y)
	}

	if robots[1].X != ArenaSize-1 || robots[1].Y != ArenaSize-1 {
		t.Errorf("Robot 2 position is incorrect. Found %d x %d, expected %d x %d.", robots[1].X, robots[1].Y, ArenaSize-1, ArenaSize-1)
	}

}

func TestNewGameLimitRobots(t *testing.T) {

	NewGame()

	_, error := JoinGame(1, 1, 0)
	if error != nil {
		t.Errorf("Tried to add the first robot but got an error: %s.", error)
	}

	_, error = JoinGame(1, 1, 0)
	if error != nil {
		t.Errorf("Tried to add a second robot but got an error: %s.", error)
	}

	_, error = JoinGame(1, 1, 0)
	if error == nil {
		t.Errorf("Game was supposed to be full but it wasn't.")
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
