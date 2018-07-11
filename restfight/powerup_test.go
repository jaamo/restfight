package restfight

import "testing"

func TestCreatePowerup(t *testing.T) {

	NewGame()
	JoinGame(1, 1, 0)
	JoinGame(1, 1, 0)

	CreatePowerupAndAddToArena(powerupTypeSpeed, 1, 1)

	if arena[1][1].Type != 3 {
		t.Errorf("No powerup on location 1x1 (wrong cell type).")
	}

	if arena[1][1].Powerup == nil {
		t.Errorf("No powerup on location 1x1 (powerup is null).")
	}

	if len(powerups) != 1 {
		t.Errorf("Powerup not added to the global list. Length was %d.", len(powerups))
	}

}
func TestAddAndRemovePowerup(t *testing.T) {

	NewGame()
	JoinGame(1, 1, 0)
	JoinGame(1, 1, 0)

	CreatePowerupAndAddToArena(powerupTypeSpeed, 1, 1)
	AddPowerupToRobot(0, 0)

	if arena[1][1].Type != 1 {
		t.Errorf("Powerup should have been removed from location 1x1. Type was %d but should have been 1.", arena[1][1].Type)
	}

	if robots[0].Powerup == nil {
		t.Errorf("Robot should have a powerup.")
	}

	if robots[0].MaxMoves != 9 {
		t.Errorf("Robot MaxMoves is incorrect. Was %d but should be 9.", robots[0].MaxMoves)
	}

	RemovePowerupFromRobot(0)

	if robots[0].Powerup != nil {
		t.Errorf("Robot should not have a powerup.")
	}

	if robots[0].MaxMoves != 4 {
		t.Errorf("Robot MaxMoves is incorrect. Was %d but should be 4.", robots[0].MaxMoves)
	}

}

func TestArenaLifetime(t *testing.T) {

	NewGame()
	JoinGame(1, 1, 0)
	JoinGame(1, 1, 0)

	CreatePowerupAndAddToArena(powerupTypeSpeed, 1, 1)

	if powerups[0].ArenaLifetime != 5 {
		t.Errorf("Initial lifetime should be 5, was %d.", powerups[0].ArenaLifetime)
	}

	ConsumeArenaPowerups()
	if powerups[0].ArenaLifetime != 4 {
		t.Errorf("Lifetime after 1 turn should be 4, was %d.", powerups[0].ArenaLifetime)
	}

	ConsumeArenaPowerups()
	if powerups[0].ArenaLifetime != 3 {
		t.Errorf("Lifetime after 2 turns should be 3, was %d.", powerups[0].ArenaLifetime)
	}

	ConsumeArenaPowerups()
	if powerups[0].ArenaLifetime != 2 {
		t.Errorf("Lifetime after 3 turn should be 2, was %d.", powerups[0].ArenaLifetime)
	}

	ConsumeArenaPowerups()
	if powerups[0].ArenaLifetime != 1 {
		t.Errorf("Lifetime after 4 turn should be 1, was %d.", powerups[0].ArenaLifetime)
	}

	ConsumeArenaPowerups()
	if len(powerups) != 0 {
		t.Errorf("Expired powerup not removed.")
	}

}

func TestRobotLifetime(t *testing.T) {

	NewGame()
	JoinGame(1, 1, 0)
	JoinGame(1, 1, 0)

	CreatePowerupAndAddToArena(powerupTypeSpeed, 1, 1)
	AddPowerupToRobot(0, 0)

	if robots[0].Powerup.RobotLifetime != 5 {
		t.Errorf("Lifetime after 0 turns should be 5, was %d.", robots[0].Powerup.RobotLifetime)
	}
	if robots[0].MaxMoves != 9 {
		t.Errorf("Robot MaxMoves is incorrect. Was %d but should be 9.", robots[0].MaxMoves)
	}

	ConsumeRobotPowerups()
	if robots[0].Powerup.RobotLifetime != 4 {
		t.Errorf("Lifetime after 1 turns should be 4, was %d.", robots[0].Powerup.RobotLifetime)
	}

	ConsumeRobotPowerups()
	if robots[0].Powerup.RobotLifetime != 3 {
		t.Errorf("Lifetime after 2 turns should be 3, was %d.", robots[0].Powerup.RobotLifetime)
	}

	ConsumeRobotPowerups()
	if robots[0].Powerup.RobotLifetime != 2 {
		t.Errorf("Lifetime after 3 turns should be 2, was %d.", robots[0].Powerup.RobotLifetime)
	}

	ConsumeRobotPowerups()
	if robots[0].Powerup.RobotLifetime != 1 {
		t.Errorf("Lifetime after 4 turns should be 1, was %d.", robots[0].Powerup.RobotLifetime)
	}

	ConsumeRobotPowerups()
	if robots[0].Powerup != nil {
		t.Errorf("Powerup should have been expired.")
	}
	if robots[0].MaxMoves != 4 {
		t.Errorf("Robot MaxMoves is incorrect. Was %d but should be 4.", robots[0].MaxMoves)
	}

}

func TestPowerupApply(t *testing.T) {

	NewGame()
	JoinGame(1, 1, 0)
	JoinGame(1, 1, 0)

	CreatePowerupAndAddToArena(powerupTypeSpeed, 1, 1)

	if len(powerups) != 1 {
		t.Errorf("Powerup should have been added.")
	}

	DeployPowerupsFromArenaToRobots()

	if len(powerups) != 1 {
		t.Errorf("Powerup should have not been deployed.")
	}

	CreatePowerupAndAddToArena(powerupTypeSpeed, 0, 0)

	if len(powerups) != 2 {
		t.Errorf("Powerup should have been added.")
	}

	DeployPowerupsFromArenaToRobots()

	if len(powerups) != 1 {
		t.Errorf("Powerup should have been deployed to a robot.")
	}

	if robots[0].Powerup == nil {
		t.Errorf("Powerup not added to a robot.")
	}

}
