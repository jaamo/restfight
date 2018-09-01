package restfight

import (
	"fmt"
	"math/rand"
)

// PowerupRandomizer creates new powerups to the map. Call this on each turn.
func PowerupRandomizer(forceAdd bool) {

	rand.Seed(42)
	if forceAdd || rand.Intn(5) == 0 {

		fmt.Println("Add powerup!")

		var powerupTypeInt = 0 // rand.Intn(2)

		// Find all empty cells.
		var emptyCells []*Cell
		for x := 0; x < len(arena); x++ {
			for y := 0; y < len(arena); y++ {
				if arena[x][y].Type == 0 {
					emptyCells = append(emptyCells, &arena[x][y])
				}
			}
		}

		// Abort of the map is empty for some reason.
		if len(emptyCells) == 0 {
			fmt.Println("WFT?! No empty cells found.")
			return
		}

		// Randomly pick one.
		var randomCell = emptyCells[rand.Intn(len(emptyCells))]

		var powerupType = ""
		if powerupTypeInt == 0 {
			powerupType = powerupTypeSpeed
		}

		CreatePowerupAndAddToArena(powerupType, randomCell.X, randomCell.Y)

	} else {
		fmt.Println("Not now...")
	}
}

// CreatePowerupAndAddToArena adds powerup to the arena with given type.
func CreatePowerupAndAddToArena(powerupType string, x int, y int) {

	var powerup = Powerup{
		Type:               powerupType,
		X:                  x,
		Y:                  y,
		AdditionalMaxMoves: 5,
		ArenaLifetime:      5,
		RobotLifetime:      5,
	}

	arena[x][y].Type = 3
	arena[x][y].Powerup = &powerup

	powerups = append(powerups, &powerup)

}

// AddPowerupToRobot adds given powerup to robot.
func AddPowerupToRobot(robotID int, powerupID int) {

	// First of all remove current powerup.
	RemovePowerupFromRobot(robotID)

	// Then set new powerup.
	robots[robotID].Powerup = powerups[powerupID]

	// Then apply new powerup attributes.
	if powerups[powerupID].Type == powerupTypeSpeed {

		robots[robotID].MaxMoves += powerups[powerupID].AdditionalMaxMoves

	}

	// Remove powerup from the arena.
	arena[powerups[powerupID].X][powerups[powerupID].Y].Type = 1
	arena[powerups[powerupID].X][powerups[powerupID].Y].Powerup = nil

	// Remove also from popups list.
	powerups = append(powerups[:powerupID], powerups[powerupID+1:]...)

}

// RemovePowerupFromRobot removes current powerup from a given robot.
func RemovePowerupFromRobot(robotID int) {

	// Check if powerup is set.
	if robots[robotID].Powerup != nil {

		// Remove speed.
		if robots[robotID].Powerup.Type == powerupTypeSpeed {
			robots[robotID].MaxMoves -= robots[robotID].Powerup.AdditionalMaxMoves
		}

		// Reset.
		robots[robotID].Powerup = nil

	}

}

// DeployPowerupsFromArenaToRobots checks if any of the robots are on top of a powerup and applies them. Should be called after movement.
func DeployPowerupsFromArenaToRobots() {

	// Check each robot...
	for i := 0; i < len(robots); i++ {

		// ...against each powerup.
		for k := 0; k < len(powerups); k++ {

			// Robot is on a powerup!
			if powerups[k].X == robots[i].X && powerups[k].Y == robots[i].Y {

				// Add powerup to the robot.
				AddPowerupToRobot(i, k)

			}

		}

	}

}

// ConsumeArenaPowerups cleans expired powerups from the arena.
func ConsumeArenaPowerups() {

	var newPowerups []*Powerup

	for i := 0; i < len(powerups); i++ {

		// Reduce lifetime.
		powerups[i].ArenaLifetime--

		if powerups[i].ArenaLifetime > 0 {

			newPowerups = append(newPowerups, powerups[i])

		} else {

			// Remove from arena.
			arena[powerups[i].X][powerups[i].Y].Type = 0
			arena[powerups[i].X][powerups[i].Y].Powerup = nil

		}

	}

	powerups = newPowerups

}

// ConsumeRobotPowerups consumes powerups on all robots and removes expired powerups. This should be called after the end of each turn.
func ConsumeRobotPowerups() {

	// Check all robots.
	for i := 0; i < len(robots); i++ {

		// Powerup exists.
		if robots[i].Powerup != nil {

			// Reduce lifetime.
			robots[i].Powerup.RobotLifetime--

			// Expires => remove.
			if robots[i].Powerup.RobotLifetime == 0 {
				RemovePowerupFromRobot(i)
			}

		}

	}

}
