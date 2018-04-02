package restfight

import "testing"

func TestCanPlay(t *testing.T) {

	NewGame()
	JoinGame()
	JoinGame()

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
	JoinGame()
	JoinGame()

	if turn != 0 {
		t.Errorf("New game. Turn should be 0.")
	}

	ToggleTurn()

	if turn != 1 {
		t.Errorf("2nd turn. Turn should be 1.")
	}

	ToggleTurn()

	if turn != 0 {
		t.Errorf("3rd turn. Turn should be 0.")
	}

}
