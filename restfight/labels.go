package restfight

// Powerup type labels.
const powerupTypeSpeed = "POWERUP_SPEED"

/**
 * Datatype defining game status.
 */
type GameStatus int

/**
 * Game status constannts.
 */
const (
	GameStatusWaitingForPlayers GameStatus = iota
	GameStatusRunning
	GameStatusGameOver
)

/**
 * Labels for game statuses.
 */
var GameStatusLabels = [...]string{
	"WAITING_FOR_PLAYERS",
	"DEPLOYMENT",
	"RUNNING",
	"GAME_OVER",
}

/**
 * Active robot status.
 */
type ActiveRobotStatus int

/**
 * Game status constannts.
 */
const (
	ActiveRobotStatusWaiting ActiveRobotStatus = iota
	ActiveRobotStatusPlaying
)

/**
 * Labels for game statuses.
 */
var ActiveRobotStatusLabels = [...]string{
	"WAITING",
	"PLAYING",
}
