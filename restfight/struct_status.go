package restfight

/**
 * Status data model for the game.
 */
type Status struct {

	// Unique game id.
	GameID int `json:"game_id,omitempty"`

	// Game status: 0 = waiting for players, 1 = robot deployment, 2 = game is on, 3 = game over
	Status GameStatus `json:"status,omitempty"`

	// Active robot. 0 or 1.
	ActiveRobot int `json:"active_robot,omitempty"`

	// Active robot status. 0 = waiting, 1 = turn started
	ActiveRobotStatus ActiveRobotStatus `json:"active_robot_status,omitempty"`
}

/**
 * Datatype defining game status.
 */
type GameStatus int

/**
 * Game status constannts.
 */
const (
	GameStatusWaitingForPlayers GameStatus = iota
	GameStatusDeployment
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
