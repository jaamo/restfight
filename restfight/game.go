package restfight

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

// NewGame starts new game.
func NewGame() {

	// Init empty arena.
	for x := 0; x < ArenaSize; x++ {
		for y := 0; y < ArenaSize; y++ {
			arena[x][y] = Cell{Type: ArenaTypeEmpty, X: x, Y: y}
		}
	}

	randomMap(&arena)

	// Init game status
	status = Status{
		GameID:      generateKey(1, 100),
		Status:      GameStatusWaitingForPlayers,
		ActiveRobot: 0,
		Arena:       &arena,
		Robots:      &robots,
		Powerups:    &powerups,
	}

	// Reset turn.
	status.ActiveRobot = 0
	status.IsYourTurn = 0

	// Clear all players.
	robots = robots[:0]
	powerups = powerups[:0]

}

// JoinGame add a new robot with specified info to the arena and return it. Return an error if game is full.
func JoinGame(engineLevel int, shieldLevel int, weaponLevel int) (Robot, error) {

	var robot Robot
	var err error

	if len(robots) == 2 {
		return robot, errors.New("GAME_FULL")
	}

	// Robot coordinates.
	x := 0
	y := 0
	if len(robots) == 1 {
		x = ArenaSize - 1
		y = ArenaSize - 1
		// x = 0
		// y = 1
	}

	// Create robot.
	robot, err = CreateRobot(engineLevel, shieldLevel, weaponLevel)
	if err != nil {
		return robot, err
	}
	robot.RobotID = generateKey(len(robots), 100)
	robot.X = x
	robot.Y = y
	robot.RobotIndex = len(robots)
	robots = append(robots, &robot)

	// Two players joined, set turn.
	if len(robots) == 2 {
		status.Status = GameStatusRunning
	}

	// Add to arena.
	arena[x][y].Type = ArenaTypeRobot
	arena[x][y].Robot = &robot

	return robot, nil

}

// CanPlay takes robot index as an argument and return true if the given robot is active (is that robot's turn)
func CanPlay(robotIndex int) bool {
	return status.ActiveRobot == robotIndex
}

// ToggleTurn switches turn to another robot.
func ToggleTurn() {

	// Swap turn.
	status.ActiveRobot = (status.ActiveRobot + 1) % 2

	// Regeneration. Add health to an active robot.
	if robots[status.ActiveRobot].Health < robots[status.ActiveRobot].MaxHealth {
		robots[status.ActiveRobot].Health++
	}

	// Reset ammo and moves.
	for i := 0; i < len(robots); i++ {
		robots[i].WeaponAmmo = 1
		robots[i].Moves = robots[i].MaxMoves
	}

	// When beginning of a new round consume powerups and randomly add new poweups.
	if status.ActiveRobot == 0 {
		fmt.Println("Add powerup?")
		PowerupRandomizer(false)
		ConsumeArenaPowerups()
		ConsumeRobotPowerups()
		// fmt.Printf("%v", powerups)
	}

}

// GetStatus is only debugging atm.
func GetStatus(robotIndex int) Status {

	//var arena [ArenaSize][ArenaSize]Cell

	var output string
	for x := 0; x < ArenaSize; x++ {
		for y := 0; y < ArenaSize; y++ {
			output += strconv.Itoa(arena[x][y].Type) + " "
		}
		output += "\n"
	}

	status.Robot = nil
	status.Enemies = nil
	status.IsYourTurn = 0

	// Start new game if coundown is zero.
	if status.Status == GameStatusGameOver {
		gameOverDuration := time.Now().Unix() - gameOverTime
		fmt.Printf("Game over duration: %d / 10 seconds.\n", gameOverDuration)
		if gameOverDuration > 10 {
			NewGame()
		}
	}

	// No robot defined so skip the rest.
	if robotIndex < 0 {
		return status
	}

	// Get players robot.
	if len(robots) > 0 && robotIndex < len(robots) {

		status.Robot = robots[robotIndex]

		if status.Status == GameStatusRunning && robotIndex == status.ActiveRobot {
			status.IsYourTurn = 1
		}

	}

	// Add enemy robot to list.
	if len(robots) > 1 {
		var enemies []*Robot
		enemies = append(enemies, robots[((robotIndex+1)%2)])
		status.Enemies = enemies
	}

	return status
}

// updateStatus updates status based on robots alive.
func updateStatus() {

	for i := 0; i < len(robots); i++ {
		if robots[i].Health <= 0 {
			status.Status = GameStatusGameOver
			gameOverTime = time.Now().Unix()
		}
	}

}
